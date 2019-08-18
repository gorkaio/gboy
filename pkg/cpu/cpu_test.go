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
	mem.EXPECT().Read(PCStartAddress).Return(byte(0xDB))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(0x00))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(0x00))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(0x00))

	c := cpu.New(mem)
	_, err := c.Step()
	assert.Error(t, err, "Unknown opcode 0xDB")
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
	c.SetFlagZ(true)
	assert.True(t, c.FlagZ())
}
func TestClearsZFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0xFF)
	c.SetFlagZ(false)
	assert.False(t, c.FlagZ())
}

func TestSetsCFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0x00)
	c.SetFlagC(true)
	assert.True(t, c.FlagC())
}

func TestClearsCFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0xFF)
	c.SetFlagC(false)
	assert.False(t, c.FlagC())
}

func TestSetsHFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0x00)
	c.SetFlagH(true)
	assert.True(t, c.FlagH())
}
func TestClearsHFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0xFF)
	c.SetFlagH(false)
	assert.False(t, c.FlagH())
}

func TestSetsNFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0x00)
	c.SetFlagN(true)
	assert.True(t, c.FlagN())
}
func TestClearsNFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)
	c.F.Set(0xFF)
	c.SetFlagN(false)
	assert.False(t, c.FlagN())
}
