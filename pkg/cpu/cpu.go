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
	Registers Registers
	memory    memory.Controller
}

// New initialises a new Z80 cpu
func New(memory memory.Controller) *CPU {
	return &CPU{
		Registers: Registers{
			PC: uint16(0x100),
		},
		memory: memory,
	}
}

// Step executes next instruction and returns cycles consumed
func (cpu *CPU) Step() int {
	if cpu.memory.Read(cpu.Registers.PC) == 0x00 {
		cpu.Registers.PC++
		return 4
	}
	panic("Unknown OpCode")
}
