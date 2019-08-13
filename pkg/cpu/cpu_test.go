package cpu

import (
	gomock "github.com/golang/mock/gomock"
	memory "github.com/gorkaio/gboy/pkg/memory"
	assert "gotest.tools/assert"
	testing "testing"
)

const PCStartAddress = uint16(0x100)
const NOP = byte(0x00)
const NOPCycles = 4

func TestProgramCounterStartsAt0x100(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := memory.NewMockController(ctrl)
	cpu := NewCPU(mem)
	assert.Equal(t, cpu.registers.PC, uint16(0x100))
}

func TestExecutesInstructions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mem := memory.NewMockController(ctrl)
	mem.EXPECT().Read(PCStartAddress).Return(NOP)

	cpu := NewCPU(mem)
	pc := cpu.registers.PC

	assert.Equal(t, cpu.Step(), NOPCycles)
	assert.Equal(t, true, cpu.registers.PC > pc)
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

	cpu := NewCPU(mem)
	cpu.Step()
}