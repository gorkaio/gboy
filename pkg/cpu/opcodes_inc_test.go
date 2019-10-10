package cpu_test

import (
	"fmt"
	"testing"
)

func TestIncrementClearsNegativeFlagIn8BitRegisters(t *testing.T) {
	for r8, opc := range incOpcodesFor8bitRegisters() {
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
	}
}

func TestIncrementSetsZeroFlagWithoutAffectingCarryIn8BitRegisters(t *testing.T) {
	for r8, opc := range incOpcodesFor8bitRegisters() {
		testFlagZ := testDescription{
			fmt.Sprintf("'INC %s' sets zero flag when result is zero without affecting carry flag", r8),
			opc,
			regMap{r8: 0xFF},
			regMap{r8: 0x00, "F": FlagZ | FlagH},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagZ)
		testCase.Run(t)
	}
}

func TestIncrementDoesNotClearCarryFlagWhenSetIn8BitRegisters(t *testing.T) {
	for r8, opc := range incOpcodesFor8bitRegisters() {
		testFlagC := testDescription{
			fmt.Sprintf("'INC %s' does not clear carry flag when it was previously set", r8),
			opc,
			regMap{r8: 0x12, "F": FlagC},
			regMap{r8: 0x13, "F": FlagC},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagC)
		testCase.Run(t)
	}
}

func TestIncrementSetHalfCarryFor8BitRegisters(t *testing.T) {
	for r8, opc := range incOpcodesFor8bitRegisters() {
		testFlagH := testDescription{
			fmt.Sprintf("'INC %s' sets half-carry flag when increment carries in bits 3-4", r8),
			opc,
			regMap{r8: 0x0F},
			regMap{r8: 0x10, "F": FlagH},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagH)
		testCase.Run(t)
	}
}

func TestIncrementClearsNegativeFlagIn16bitRegisters(t *testing.T) {
	for r16, opc := range incOpcodesFor16bitRegisters() {
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
	}
}

func TestIncrementSetsZeroFlagWithoutAffectingCarryIn16bitRegisters(t *testing.T) {
	for r16, opc := range incOpcodesFor16bitRegisters() {
		testFlagZ := testDescription{
			fmt.Sprintf("'INC %s' sets zero flag when result is zero without affecting carry flag", r16),
			opc,
			regMap{r16: 0xFFFF},
			regMap{r16: 0x0000, "F": FlagZ | FlagH},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagZ)
		testCase.Run(t)
	}
}

func TestIncrementDoesNotClearCarryFlagWhenSetIn16bitRegisters(t *testing.T) {
	for r16, opc := range incOpcodesFor16bitRegisters() {
		testFlagC := testDescription{
			fmt.Sprintf("'INC %s' does not clear carry flag when it was previously set", r16),
			opc,
			regMap{r16: 0x1234, "F": FlagC},
			regMap{r16: 0x1235, "F": FlagC},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagC)
		testCase.Run(t)
	}
}

func TestIncrementSetHalfCarryFor16bitRegisters(t *testing.T) {
	for r16, opc := range incOpcodesFor16bitRegisters() {
		testFlagH := testDescription{
			fmt.Sprintf("'INC %s' sets half-carry flag when increment carries in bits 11-12", r16),
			opc,
			regMap{r16: 0x0FFF},
			regMap{r16: 0x1000, "F": FlagH},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagH)
		testCase.Run(t)
	}
}

func incOpcodesFor8bitRegisters() map[string]opcode {
	return map[string]opcode{
		"A": {0x3C},
		"B": {0x04}, "C": {0x0C},
		"D": {0x14}, "E": {0x1C},
		"H": {0x24}, "L": {0x2C},
	}
}

func incOpcodesFor16bitRegisters() map[string]opcode {
	return map[string]opcode{
		"BC": {0x03}, "DE": {0x13},
		"HL": {0x23}, "SP": {0x33},
	}
}
