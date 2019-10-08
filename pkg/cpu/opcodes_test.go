package cpu_test

import (
	"github.com/golang/mock/gomock"
	"github.com/gorkaio/gboy/pkg/cpu"
	mocks "github.com/gorkaio/gboy/pkg/cpu/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FlagZ = 0x80
	FlagN = 0x40
	FlagH = 0x20
	FlagC = 0x10
)

type memoryAccess struct {
	address uint16
	data uint8
}

type testCase struct {
	description string
	instruction []byte
	initialState, expectedState cpu.State
	expectedReads, expectedWrites []memoryAccess
	expectedCycles int
}

func testInstruction(t *testing.T, test testCase) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)
	readInstruction(mem, test.initialState.PC, test.instruction)
	setExpectedReads(mem, test.expectedReads)
	setExpectedWrites(mem, test.expectedWrites)

	c.SetStatus(test.initialState)
	cycles, err := c.Step()
	assert.NoError(t, err, "CPU execution failed for test \"%s\"", test.description)
	assert.Equal(t, cycles, test.expectedCycles, "Expected CPU cycles do not match for test \"%s\"", test.description)

	finalStatus := c.Status()
	assert.Equal(t, test.expectedState, finalStatus, "Expected state does not match in test \"%s\"", test.description)
}

func readInstruction(mem *mocks.MockMemory, address uint16, instruction []byte) {
	for i, b := range instruction {
		mem.EXPECT().Read(address + uint16(i)).Return(b)
	}
}

func setExpectedReads(mem *mocks.MockMemory, reads []memoryAccess) {
	if (len(reads) == 0) {
		mem.EXPECT().Read(gomock.Any()).Times(0)
	}

	for _, read := range(reads) {
		mem.EXPECT().Read(read.address).Return(read.data)
	}
}

func setExpectedWrites(mem *mocks.MockMemory, writes []memoryAccess) {
	if (len(writes) == 0) {
		mem.EXPECT().Write(gomock.Any(), gomock.Any()).Times(0)
	}

	for _, write := range(writes) {
		mem.EXPECT().Write(write.address, write.data)
	}
}

