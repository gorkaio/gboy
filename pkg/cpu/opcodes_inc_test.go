package cpu_test

import (
	"fmt"
	"testing"
)

func TestIncrementFor8BitRegistersClearsNegativeFlag(t *testing.T) {
	registerOpcodes := map[string]opcode{
		"A": opcode{0x3C},
		"B": opcode{0x04}, "C": opcode{0x0C},
		"D": opcode{0x14}, "E": opcode{0x1C},
		"H": opcode{0x24}, "L": opcode{0x2C},
	}

	for r8, opc := range registerOpcodes {
		testFlagN := testDescription{
			fmt.Sprintf("'INC %s' increments %s register and clears negative flag", r8, r8),
			opc,
			regMap{r8: 0x12, "F": FlagN},
			regMap{r8: 0x13, "F": 0},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)

		testFlagZ := testDescription{
			fmt.Sprintf("'INC %s' sets zero flag when result is zero without affecting carry flag", r8),
			opc,
			regMap{r8: 0xFF},
			regMap{r8: 0x00, "F": FlagZ | FlagH},
			memMap{},
			memMap{},
			4,
		}
		testCase = buildTestCase(testFlagZ)
		testCase.Run(t)

		testFlagC := testDescription{
			fmt.Sprintf("'INC %s' does not clear carry flag when it was previously set", r8),
			opc,
			regMap{r8: 0x12, "F": FlagC},
			regMap{r8: 0x13, "F": FlagC},
			memMap{},
			memMap{},
			4,
		}
		testCase = buildTestCase(testFlagC)
		testCase.Run(t)

		testFlagH := testDescription{
			fmt.Sprintf("'INC %s' sets half-carry flag when increment carries in bits 7-8", r8),
			opc,
			regMap{r8: 0x0F},
			regMap{r8: 0x10, "F": FlagH},
			memMap{},
			memMap{},
			4,
		}
		testCase = buildTestCase(testFlagH)
		testCase.Run(t)
	}
}

func TestIncrementFor16BitRegisters(t *testing.T) {
	registerOpcodes := map[string]opcode{
		"BC": opcode{0x03}, "DE": opcode{0x13},
		"HL": opcode{0x23}, "SP": opcode{0x33},
	}

	for r16, opc := range registerOpcodes {
		testFlagN := testDescription{
			fmt.Sprintf("'INC %s' increments %s register and clears negative flag", r16, r16),
			opc,
			regMap{r16: 0x1234, "F": FlagN},
			regMap{r16: 0x1235, "F": 0},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)

		testFlagZ := testDescription{
			fmt.Sprintf("'INC %s' sets zero flag when result is zero without affecting carry flag", r16),
			opc,
			regMap{r16: 0xFFFF},
			regMap{r16: 0x0000, "F": FlagZ | FlagH},
			memMap{},
			memMap{},
			8,
		}
		testCase = buildTestCase(testFlagZ)
		testCase.Run(t)

		testFlagC := testDescription{
			fmt.Sprintf("'INC %s' does not clear carry flag when it was previously set", r16),
			opc,
			regMap{r16: 0x1234, "F": FlagC},
			regMap{r16: 0x1235, "F": FlagC},
			memMap{},
			memMap{},
			8,
		}
		testCase = buildTestCase(testFlagC)
		testCase.Run(t)

		testFlagH := testDescription{
			fmt.Sprintf("'INC %s' sets half-carry flag when increment carries in bits 7-8", r16),
			opc,
			regMap{r16: 0xF0FF},
			regMap{r16: 0xF100, "F": FlagH},
			memMap{},
			memMap{},
			8,
		}
		testCase = buildTestCase(testFlagH)
		testCase.Run(t)
	}
}
