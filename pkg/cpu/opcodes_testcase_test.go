package cpu_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorkaio/gboy/pkg/cpu"
	mocks "github.com/gorkaio/gboy/pkg/cpu/mocks"
	"github.com/stretchr/testify/assert"
)

type testState struct {
	regs map[string]uint16
	ime  bool
}

type memoryAccess struct {
	address uint16
	data    byte
}

type testCase struct {
	description                   string
	instruction                   []byte
	initialState, expectedState   testState
	expectedReads, expectedWrites []memoryAccess
	expectedCycles                int
}

func NewTestState() testState {
	return testState{
		regs: map[string]uint16{
			"AF": 0,
			"BC": 0,
			"DE": 0,
			"HL": 0,
			"SP": 0,
			"PC": 0,
		},
		ime: false,
	}
}

func NewTestCase(instruction []byte) testCase {
	return testCase{
		description:    "",
		instruction:    instruction,
		initialState:   NewTestState(),
		expectedState:  NewTestState(),
		expectedReads:  []memoryAccess{},
		expectedWrites: []memoryAccess{},
		expectedCycles: 0,
	}
}

func NewTestStateFromCPU(c cpu.State) testState {
	return testState{
		regs: map[string]uint16{
			"AF": c.AF,
			"BC": c.BC,
			"DE": c.DE,
			"HL": c.HL,
			"SP": c.SP,
			"PC": c.PC,
		},
		ime: c.IME,
	}
}

func (t *testState) toCPUState() cpu.State {
	return cpu.State{
		AF:  t.regs["AF"],
		BC:  t.regs["BC"],
		DE:  t.regs["DE"],
		HL:  t.regs["HL"],
		SP:  t.regs["SP"],
		PC:  t.regs["PC"],
		IME: t.ime,
	}
}

func (t *testState) SetRegister(regName string, data int) {
	switch regName {
	case "AF":
		t.regs["AF"] = uint16(data)
	case "BC":
		t.regs["BC"] = uint16(data)
	case "DE":
		t.regs["DE"] = uint16(data)
	case "HL":
		t.regs["HL"] = uint16(data)
	case "SP":
		t.regs["SP"] = uint16(data)
	case "PC":
		t.regs["PC"] = uint16(data)
	case "A":
		t.regs["AF"] = setHighByte(t.regs["AF"], uint8(data))
	case "F":
		t.regs["AF"] = setLowByte(t.regs["AF"], uint8(data))
	case "B":
		t.regs["BC"] = setHighByte(t.regs["BC"], uint8(data))
	case "C":
		t.regs["BC"] = setLowByte(t.regs["BC"], uint8(data))
	case "D":
		t.regs["DE"] = setHighByte(t.regs["DE"], uint8(data))
	case "E":
		t.regs["DE"] = setLowByte(t.regs["DE"], uint8(data))
	case "H":
		t.regs["HL"] = setHighByte(t.regs["HL"], uint8(data))
	case "L":
		t.regs["HL"] = setLowByte(t.regs["HL"], uint8(data))
	}
}

func setHighByte(word uint16, data uint8) uint16 {
	return (word & 0x00FF) | (uint16(data) << 8)
}

func setLowByte(word uint16, data uint8) uint16 {
	return (word & 0xFF00) | uint16(data)
}

func (t *testCase) WithRegister(regName string, data int) {
	t.initialState.SetRegister(regName, data)
	t.expectedState.SetRegister(regName, data)
}

func (t *testCase) ExpectRegister(regName string, data int) {
	t.expectedState.SetRegister(regName, data)
}

func (t *testCase) WithIME(ime bool) {
	t.initialState.ime = ime
	t.expectedState.ime = ime
}

func (t *testCase) WithDescription(description string) {
	t.description = description
}

func (t *testCase) ExpectIME(ime bool) {
	t.expectedState.ime = ime
}

func (t *testCase) ExpectCycles(cycles int) {
	t.expectedCycles = cycles
}

func setExpectedReads(mem *mocks.MockMemory, reads []memoryAccess) {
	if len(reads) == 0 {
		mem.EXPECT().Read(gomock.Any()).Times(0)
	}

	for _, read := range reads {
		mem.EXPECT().Read(read.address).Return(read.data)
	}
}

func setExpectedWrites(mem *mocks.MockMemory, writes []memoryAccess) {
	if len(writes) == 0 {
		mem.EXPECT().Write(gomock.Any(), gomock.Any()).Times(0)
	}

	for _, write := range writes {
		mem.EXPECT().Write(write.address, write.data)
	}
}

func (t *testCase) ExpectMemoryRead(address uint16, data uint8) {
	t.expectedReads = append(t.expectedReads, memoryAccess{address: address, data: data})
}

func (t *testCase) ExpectMemoryWrite(address uint16, data uint8) {
	t.expectedWrites = append(t.expectedWrites, memoryAccess{address: address, data: data})
}

func (test *testCase) Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)
	readInstruction(mem, uint16(test.initialState.regs["PC"]), test.instruction)
	setExpectedReads(mem, test.expectedReads)
	setExpectedWrites(mem, test.expectedWrites)

	c.SetStatus(test.initialState.toCPUState())

	cycles, err := c.Step()
	assert.NoError(t, err, "CPU execution failed for test case \"%s\"", test.description)
	if test.expectedCycles > 0 {
		assert.Equal(t, test.expectedCycles, cycles, "Expected CPU cycles do not match for test case \"%s\"", test.description)
	}

	finalStatus := NewTestStateFromCPU(c.Status())
	assert.Equal(t, test.expectedState, finalStatus, "Expected state does not match in test case \"%s\"", test.description)
}
