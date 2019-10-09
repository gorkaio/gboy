package cpu_test

import (
	"fmt"
	"testing"
)

func TestLoadImmediateFor8BitRegisters(t *testing.T) {
	registerOpcodes := map[string]opcode{
		"A": opcode{0x3E, 0x12},
		"B": opcode{0x06, 0x12}, "C": opcode{0x0E, 0x12},
		"D": opcode{0x16, 0x12}, "E": opcode{0x1E, 0x12},
		"H": opcode{0x26, 0x12}, "L": opcode{0x2E, 0x12},
	}

	for r8, opc := range registerOpcodes {
		testDescription := testDescription{
			fmt.Sprintf("'LD %s, %#02x' loads %#02x value into %s", r8, 0x12, 0x12, r8),
			opc,
			regMap{},
			regMap{r8: 0x12},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestLoadImmediateFor16BitRegisters(t *testing.T) {
	registerOpcodes := map[string]opcode{
		"BC": opcode{0x01, 0x34, 0x12}, "DE": opcode{0x11, 0x34, 0x12},
		"HL": opcode{0x21, 0x34, 0x12}, "SP": opcode{0x31, 0x34, 0x12},
	}

	for r16, opc := range registerOpcodes {
		testDescription := testDescription{
			fmt.Sprintf("'LD %s, %#04x' loads %#04x value into %s", r16, 0x1234, 0x1234, r16),
			opc,
			regMap{},
			regMap{r16: 0x1234},
			memMap{},
			memMap{},
			12,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}
