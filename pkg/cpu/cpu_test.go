package cpu_test

import (
	"github.com/golang/mock/gomock"
	"github.com/gorkaio/gboy/pkg/cpu"
	mocks "github.com/gorkaio/gboy/pkg/cpu/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const PCStartAddress = uint16(0x100)

func TestProgramCounterStartsAt0x100(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	c := cpu.New(mem)
	assert.Equal(t, uint16(0x100), c.Status()["PC"])
}

func TestExecutesInstructions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(cpu.NOP))

	c := cpu.New(mem)
	c.DebugDisable()
	pc := c.Status()["PC"].(uint16)
	cyclesConsumed, err := c.Step()
	assert.NoError(t, err)

	assert.True(t, cyclesConsumed > 0)
	assert.True(t, c.Status()["PC"].(uint16) > pc)
}

func TestErrorsWithUnknownOpcodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemory(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(byte(0xDB))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(0x00))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(0x00))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(0x00))

	c := cpu.New(mem)
	c.DebugDisable()
	_, err := c.Step()
	assert.Error(t, err, "Unknown opcode 0xDB")
}

func TestExecutesNOP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemory(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(uint8(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 1).Return(uint8(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 2).Return(uint8(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 3).Return(uint8(cpu.NOP))

	c := cpu.New(mem)
	c.DebugDisable()
	cycles, err := c.Step()
	assert.NoError(t, err)
	assert.Equal(t, cycles, 4)
}