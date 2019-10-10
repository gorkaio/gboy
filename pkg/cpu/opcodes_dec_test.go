package cpu_test

import (
	"fmt"
	"testing"
)

func TestDecSetsNegativeFlagIn8bitRegisters(t *testing.T) {
	for r8, opc := range decOpcodesFor8bitRegisters() {
		testFlagN := testDescription{
			fmt.Sprintf("'DEC %s' decrements %s register and set negative flag", r8, r8),
			opc,
			regMap{r8: 0x12},
			regMap{r8: 0x11, "F": FlagN},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestDecrementSetsNegativeFlagIn16bitRegisters(t *testing.T) {
	for r16, opc := range decOpcodesFor16bitRegisters() {
		testFlagN := testDescription{
			fmt.Sprintf("'DEC %s' decrements %s register and sets negative flag", r16, r16),
			opc,
			regMap{r16: 0x1234},
			regMap{r16: 0x1233, "F": FlagN},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestDecrementSetsZeroFlagWithoutAffectingCarryIn8BitRegisters(t *testing.T) {
	for r8, opc := range decOpcodesFor8bitRegisters() {
		testFlagZ := testDescription{
			fmt.Sprintf("'DEC %s' sets zero flag when result is zero without affecting carry flag", r8),
			opc,
			regMap{r8: 0x01},
			regMap{r8: 0x00, "F": FlagZ | FlagN},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagZ)
		testCase.Run(t)
	}
}

func TestDecrementSetsZeroFlagWithoutAffectingCarryIn16bitRegisters(t *testing.T) {
	for r16, opc := range decOpcodesFor16bitRegisters() {
		testFlagZ := testDescription{
			fmt.Sprintf("'DEC %s' sets zero flag when result is zero without affecting carry flag", r16),
			opc,
			regMap{r16: 0x0001},
			regMap{r16: 0x0000, "F": FlagZ | FlagN},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagZ)
		testCase.Run(t)
	}
}

func TestDecrementDoesNotClearCarryFlagWhenSetIn8BitRegisters(t *testing.T) {
	for r8, opc := range decOpcodesFor8bitRegisters() {
		testFlagC := testDescription{
			fmt.Sprintf("'DEC %s' does not clear carry flag when it was previously set", r8),
			opc,
			regMap{r8: 0x13, "F": FlagC},
			regMap{r8: 0x12, "F": FlagC | FlagN},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagC)
		testCase.Run(t)
	}
}

func TestDecrementDoesNotClearCarryFlagWhenSetIn16bitRegisters(t *testing.T) {
	for r16, opc := range decOpcodesFor16bitRegisters() {
		testFlagC := testDescription{
			fmt.Sprintf("'DEC %s' does not clear carry flag when it was previously set", r16),
			opc,
			regMap{r16: 0x1234, "F": FlagC},
			regMap{r16: 0x1233, "F": FlagC | FlagN},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagC)
		testCase.Run(t)
	}
}

func TestDecrementSetsHalfCarryFor8BitRegisters(t *testing.T) {
	for r8, opc := range decOpcodesFor8bitRegisters() {
		testFlagH := testDescription{
			fmt.Sprintf("'DEC %s' sets half-carry flag when decrement carries in bits 3-4", r8),
			opc,
			regMap{r8: 0x10},
			regMap{r8: 0x0F, "F": FlagH | FlagN},
			memMap{},
			memMap{},
			4,
		}
		testCase := buildTestCase(testFlagH)
		testCase.Run(t)
	}
}

func TestDecrementSetsHalfCarryFor16bitRegisters(t *testing.T) {
	for r16, opc := range decOpcodesFor16bitRegisters() {
		testFlagH := testDescription{
			fmt.Sprintf("'DEC %s' sets half-carry flag when decrement carries in bits 7-8", r16),
			opc,
			regMap{r16: 0x1000},
			regMap{r16: 0x0FFF, "F": FlagH | FlagN},
			memMap{},
			memMap{},
			8,
		}
		testCase := buildTestCase(testFlagH)
		testCase.Run(t)
	}
}

func decOpcodesFor8bitRegisters() map[string]opcode {
	return map[string]opcode{
		"A": {0x3D},
		"B": {0x05}, "C": {0x0D},
		"D": {0x15}, "E": {0x1D},
		"H": {0x25}, "L": {0x2D},
	}
}

func decOpcodesFor16bitRegisters() map[string]opcode {
	return map[string]opcode{
		"BC": {0x0B}, "DE": {0x1B},
		"HL": {0x2B}, "SP": {0x3B},
	}
}
