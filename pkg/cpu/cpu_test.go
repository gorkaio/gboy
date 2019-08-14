package cpu_test

import (
	gomock "github.com/golang/mock/gomock"
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	memory "github.com/gorkaio/gboy/pkg/memory/mock"
	assert "gotest.tools/assert"
	testing "testing"
)

const PCStartAddress = uint16(0x100)
const NOP = byte(0x00)

func TestProgramCounterStartsAt0x100(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := memory.NewMockController(ctrl)
	c := cpu.New(mem)
	assert.Equal(t, c.Registers.PC, uint16(0x100))
}

func TestExecutesInstructions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := memory.NewMockController(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(NOP)

	c := cpu.New(mem)
	pc := c.Registers.PC
	cyclesConsumed := c.Step()

	assert.Equal(t, true, cyclesConsumed > 0)
	assert.Equal(t, true, c.Registers.PC > pc)
}

func TestPanicsWithUnknownOpcodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic with unknown opcode")
		}
	}()

	mem := memory.NewMockController(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(byte(0xFE))

	c := cpu.New(mem)
	c.Step()
}
