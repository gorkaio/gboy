package cpu

import (
	memory "github.com/gorkaio/gboy/pkg/memory"
)

// Registers are CPU registeres
type Registers struct {
	PC uint16
}

// CPU structure
type CPU struct {
	registers Registers
	memory memory.Controller
}

// NewCPU initialises a new Z80 cpu
func NewCPU(memory memory.Controller) *CPU {
	cpu := CPU{
		registers: Registers{
			PC: uint16(0x100),
		},
		memory: memory,
	}
	return &cpu
}

// Step executes next instruction and returns cycles consumed
func (cpu *CPU) Step() int {
	if (cpu.memory.Read(cpu.registers.PC) == 0x00) {
		cpu.registers.PC++
		return 4
	}
	panic("Unknown OpCode")
}