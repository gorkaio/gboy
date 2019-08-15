package cpu_test

import (
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestDisassemblesOpCodes(t *testing.T) {
	instruction, err := cpu.Dasm(0x00)
	assert.NoError(t, err)
	assert.Equal(t, instruction, "NOP")
}

func TestFailsToDisassembleUnknownOpCodes(t *testing.T) {
	instruction, err := cpu.Dasm(0xFF)
	assert.Error(t, err, "Unknown opcode 0xFF")
	assert.Equal(t, instruction, "?")
}

func TestAssemblesInstructions(t *testing.T) {
	opcode, err := cpu.Asm("NOP")
	assert.NoError(t, err)
	assert.Equal(t, opcode, 0x00)
}

func TestFailsToAssembleUnknownInstructions(t *testing.T) {
	opcode, err := cpu.Asm("WTF")
	assert.Error(t, err)
	assert.Equal(t, opcode, 0x00)
}