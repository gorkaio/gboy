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
	mem.EXPECT().Read(PCStartAddress).Return(byte(cpu.EI))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(cpu.NOP))

	c := cpu.New(mem)
	_, err := c.Step()
	assert.NoError(t, err)
	assert.True(t, c.InterruptsEnabled())
}

func TestDIDisablesInterruptMasterEnableFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := mocks.NewMockMemoryInterface(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(byte(cpu.DI))
	mem.EXPECT().Read(PCStartAddress + 1).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 2).Return(byte(cpu.NOP))
	mem.EXPECT().Read(PCStartAddress + 3).Return(byte(cpu.NOP))

	c := cpu.New(mem)
	_, err := c.Step()
	assert.NoError(t, err)
	assert.False(t, c.InterruptsEnabled())
}
