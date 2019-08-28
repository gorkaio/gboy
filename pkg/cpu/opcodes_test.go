package cpu_test

import (
	gomock "github.com/golang/mock/gomock"
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	mocks "github.com/gorkaio/gboy/pkg/mocks"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestEIEnablesInterruptMasterEnableFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	readInstruction(mem, PCStartAddress, []byte{cpu.EI, 0x00, 0x00, 0x00})

	c := cpu.New(mem)
	_, err := c.Step()
	assert.NoError(t, err)
	assert.True(t, c.InterruptsEnabled())
}

func TestDIDisablesInterruptMasterEnableFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	readInstruction(mem, PCStartAddress, []byte{cpu.DI, 0x00, 0x00, 0x00})

	c := cpu.New(mem)
	_, err := c.Step()
	assert.NoError(t, err)
	assert.False(t, c.InterruptsEnabled())
}

func TestDECDecrementsByteRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    cpu.ByteRegisterInterface
		cycles      int
	}{
		{0x100, []byte{cpu.DEC_A, 0x00, 0x00, 0x00}, c.A, 4},
		{0x200, []byte{cpu.DEC_B, 0x00, 0x00, 0x00}, c.B, 4},
		{0x300, []byte{cpu.DEC_C, 0x00, 0x00, 0x00}, c.C, 4},
		{0x400, []byte{cpu.DEC_D, 0x00, 0x00, 0x00}, c.D, 4},
		{0x500, []byte{cpu.DEC_E, 0x00, 0x00, 0x00}, c.E, 4},
		{0x600, []byte{cpu.DEC_H, 0x00, 0x00, 0x00}, c.H, 4},
		{0x700, []byte{cpu.DEC_L, 0x00, 0x00, 0x00}, c.L, 4},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		test.register.Set(0xFF)
		c.PC.Set(test.address)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, test.register.Get(), uint8(0xFE))
		assert.Equal(t, c.PC.Get(), test.address+1)
	}
}

func TestDECDecrementsWordRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    cpu.WordRegisterInterface
		cycles      int
	}{
		{0x100, []byte{cpu.DEC_BC, 0x00, 0x00, 0x00}, c.BC, 8},
		{0x200, []byte{cpu.DEC_DE, 0x00, 0x00, 0x00}, c.DE, 8},
		{0x300, []byte{cpu.DEC_HL, 0x00, 0x00, 0x00}, c.HL, 8},
		{0x400, []byte{cpu.DEC_SP, 0x00, 0x00, 0x00}, c.SP, 8},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		test.register.Set(0xFFFF)
		c.PC.Set(test.address)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, test.register.Get(), uint16(0xFFFE))
		assert.Equal(t, c.PC.Get(), test.address+1)
	}
}

func TestINCIncrementsByteRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    cpu.ByteRegisterInterface
		cycles      int
	}{
		{0x100, []byte{cpu.INC_A, 0x00, 0x00, 0x00}, c.A, 4},
		{0x200, []byte{cpu.INC_B, 0x00, 0x00, 0x00}, c.B, 4},
		{0x300, []byte{cpu.INC_C, 0x00, 0x00, 0x00}, c.C, 4},
		{0x400, []byte{cpu.INC_D, 0x00, 0x00, 0x00}, c.D, 4},
		{0x500, []byte{cpu.INC_E, 0x00, 0x00, 0x00}, c.E, 4},
		{0x600, []byte{cpu.INC_H, 0x00, 0x00, 0x00}, c.H, 4},
		{0x700, []byte{cpu.INC_L, 0x00, 0x00, 0x00}, c.L, 4},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		test.register.Set(0xFE)
		c.PC.Set(test.address)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, test.register.Get(), uint8(0xFF))
		assert.Equal(t, c.PC.Get(), test.address+1)
	}
}

func TestINCIncrementsWordRegisters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	var tests = []struct {
		address     uint16
		instruction []byte
		register    cpu.WordRegisterInterface
		cycles      int
	}{
		{0x100, []byte{cpu.INC_BC, 0x00, 0x00, 0x00}, c.BC, 8},
		{0x200, []byte{cpu.INC_DE, 0x00, 0x00, 0x00}, c.DE, 8},
		{0x300, []byte{cpu.INC_HL, 0x00, 0x00, 0x00}, c.HL, 8},
		{0x400, []byte{cpu.INC_SP, 0x00, 0x00, 0x00}, c.SP, 8},
	}

	for _, test := range tests {
		readInstruction(mem, test.address, test.instruction)
		test.register.Set(0xFFFE)
		c.PC.Set(test.address)
		cycles, err := c.Step()

		assert.NoError(t, err)
		assert.Equal(t, cycles, test.cycles)
		assert.Equal(t, test.register.Get(), uint16(0xFFFF))
		assert.Equal(t, c.PC.Get(), test.address+1)
	}
}

func readInstruction(mem *mocks.MockMemoryInterface, address uint16, instruction []byte) {
	for i, b := range instruction {
		mem.EXPECT().Read(address + uint16(i)).Return(b)
	}
}
