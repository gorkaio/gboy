package cpu_test

import (
	"github.com/golang/mock/gomock"
	"github.com/gorkaio/gboy/pkg/cpu"
	mocks "github.com/gorkaio/gboy/pkg/cpu/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEIEnablesInterruptMasterEnableFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	readInstruction(mem, PCStartAddress, []byte{cpu.EI, 0x00, 0x00, 0x00})

	c := cpu.New(mem)
	_, err := c.Step()
	assert.NoError(t, err)
	assert.True(t, c.InterruptsEnabled())
}

func TestDIDisablesInterruptMasterEnableFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	readInstruction(mem, PCStartAddress, []byte{cpu.DI, 0x00, 0x00, 0x00})

	c := cpu.New(mem)
	_, err := c.Step()
	assert.NoError(t, err)
	assert.False(t, c.InterruptsEnabled())
}

func TestDECDecrementsByteRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    string
		expected	uint8
		cycles      int
	}{
		{0x100, []byte{cpu.DEC_A, 0x00, 0x00, 0x00}, "A", 0xFF, 4},
		{0x101, []byte{cpu.DEC_B, 0x00, 0x00, 0x00}, "B", 0xFF, 4},
		{0x102, []byte{cpu.DEC_C, 0x00, 0x00, 0x00}, "C", 0xFF, 4},
		{0x103, []byte{cpu.DEC_D, 0x00, 0x00, 0x00}, "D", 0xFF, 4},
		{0x104, []byte{cpu.DEC_E, 0x00, 0x00, 0x00}, "E", 0xFF, 4},
		{0x105, []byte{cpu.DEC_H, 0x00, 0x00, 0x00}, "H", 0xFF, 4},
		{0x106, []byte{cpu.DEC_L, 0x00, 0x00, 0x00}, "L", 0xFF, 4},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, c.Status()[test.register], test.expected)
		assert.Equal(t, c.Status()["PC"].(uint16), test.address+1)
	}
}

func TestDECDecrementsWordRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    string
		expected uint16
		cycles      int
	}{
		{0x100, []byte{cpu.DEC_BC, 0x00, 0x00, 0x00}, "BC", 0xFFFF, 8},
		{0x101, []byte{cpu.DEC_DE, 0x00, 0x00, 0x00}, "DE", 0xFFFF, 8},
		{0x102, []byte{cpu.DEC_HL, 0x00, 0x00, 0x00}, "HL", 0xFFFF, 8},
		{0x103, []byte{cpu.DEC_SP, 0x00, 0x00, 0x00}, "SP", 0xFFFF, 8},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, c.Status()[test.register], test.expected)
		assert.Equal(t, c.Status()["PC"].(uint16), test.address+1)
	}
}

func TestINCIncrementsByteRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    string
		expected uint8
		cycles      int
	}{
		{0x100, []byte{cpu.INC_A, 0x00, 0x00, 0x00}, "A", 1, 4},
		{0x101, []byte{cpu.INC_B, 0x00, 0x00, 0x00}, "B", 1, 4},
		{0x102, []byte{cpu.INC_C, 0x00, 0x00, 0x00}, "C", 1, 4},
		{0x103, []byte{cpu.INC_D, 0x00, 0x00, 0x00}, "D", 1, 4},
		{0x104, []byte{cpu.INC_E, 0x00, 0x00, 0x00}, "E", 1, 4},
		{0x105, []byte{cpu.INC_H, 0x00, 0x00, 0x00}, "H", 1, 4},
		{0x106, []byte{cpu.INC_L, 0x00, 0x00, 0x00}, "L", 1, 4},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, c.Status()[test.register], test.expected)
		assert.Equal(t, c.Status()["PC"].(uint16), test.address+1)
	}
}

func TestINCIncrementsWordRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    string
		expected uint16
		cycles      int
	}{
		{0x100, []byte{cpu.INC_BC, 0x00, 0x00, 0x00}, "BC", 1, 8},
		{0x101, []byte{cpu.INC_DE, 0x00, 0x00, 0x00}, "DE", 1, 8},
		{0x102, []byte{cpu.INC_HL, 0x00, 0x00, 0x00}, "HL", 1, 8},
		{0x103, []byte{cpu.INC_SP, 0x00, 0x00, 0x00}, "SP", 1, 8},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, c.Status()[test.register], test.expected)
		assert.Equal(t, c.Status()["PC"].(uint16), test.address+1)
	}
}

func readInstruction(mem *mocks.MockMemory, address uint16, instruction []byte) {
	for i, b := range instruction {
		mem.EXPECT().Read(address + uint16(i)).Return(b)
	}
}
