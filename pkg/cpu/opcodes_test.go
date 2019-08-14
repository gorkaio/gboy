package cpu_test

import (
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	assert "gotest.tools/assert"
	testing "testing"
)

func TestDisassemblesOpCodes(t *testing.T) {
	assert.Equal(t, cpu.Dasm(0x00), "NOP")
}

func TestAssemblesOpCodes(t *testing.T) {
	assert.Equal(t, cpu.Asm("NOP"), uint(0x00))
}

func TestPanicsDisassemblingUnknownOpCodes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic with unknown opcode")
		}
	}()

	cpu.Dasm(0xFEFFEFEF)
}

func TestPanicsAssemblingUnknownInstructions(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic with unknown opcode")
		}
	}()

	cpu.Asm("WTF")
}
