package cpu_test

import (
	gomock "github.com/golang/mock/gomock"
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	mocks "github.com/gorkaio/gboy/pkg/mocks"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

const PCStartAddress = uint16(0x100)

func TestProgramCounterStartsAt0x100(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	assert.Equal(t, uint16(0x100), c.PC.Get())
}

func TestExecutesInstructions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(cpu.NOP))

	c := cpu.New(mem)
	pc := c.PC.Get()
	cyclesConsumed, err := c.Step()
	assert.NoError(t, err)

	assert.True(t, cyclesConsumed > 0)
	assert.True(t, c.PC.Get() > pc)
}

func TestErrorsWithUnknownOpcodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(byte(0xFE))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(0xFE))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(0xFE))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(0xFE))

	c := cpu.New(mem)
	_, err := c.Step()
	assert.Error(t, err, "Unknown opcode 0xFE")
}

func TestExecutesNOP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(uint8(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 1).Return(uint8(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 2).Return(uint8(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 3).Return(uint8(cpu.NOP))

	c := cpu.New(mem)
	cycles, err := c.Step()
	assert.NoError(t, err)
	assert.Equal(t, cycles, 4)
}

func TestSetsZFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0x00)
	c.SetZ()
	assert.Equal(t, c.F.Get(), uint8(0x80))
}

func TestClearsZFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0xFF)
	c.ClearZ()
	assert.Equal(t, c.F.Get(), uint8(0x7F))
}

func TestUpdatesZFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0x00)
	c.UpdateZ(0)
	assert.Equal(t, c.F.Get(), uint8(0x80))
	c.UpdateZ(2)
	assert.Equal(t, c.F.Get(), uint8(0x00))
}

