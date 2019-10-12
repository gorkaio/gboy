package cpu_test

import (
	"fmt"
	"testing"
)

func TestAndDirectExecutesLogicalAndWithA(t *testing.T) {
	tests := []struct{
		opc opcode
		srcReg string
		init regMap
		expected regMap
		cycles int
	} {
		{opcode{0xA0}, "B", regMap{"A": 0x13, "B": 0x35, "F": FlagN | FlagC}, regMap{"A": 0x11, "F": FlagH}, 4},
		{opcode{0xA0}, "B", regMap{"A": 0x13, "B": 0x24, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
		{opcode{0xA1}, "C", regMap{"A": 0x13, "C": 0x35, "F": FlagN | FlagC}, regMap{"A": 0x11, "F": FlagH}, 4},
		{opcode{0xA1}, "C", regMap{"A": 0x13, "C": 0x24, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
		{opcode{0xA2}, "D", regMap{"A": 0x13, "D": 0x35, "F": FlagN | FlagC}, regMap{"A": 0x11, "F": FlagH}, 4},
		{opcode{0xA2}, "D", regMap{"A": 0x13, "D": 0x24, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
		{opcode{0xA3}, "E", regMap{"A": 0x13, "E": 0x35, "F": FlagN | FlagC}, regMap{"A": 0x11, "F": FlagH}, 4},
		{opcode{0xA3}, "E", regMap{"A": 0x13, "E": 0x24, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
		{opcode{0xA4}, "H", regMap{"A": 0x13, "H": 0x35, "F": FlagN | FlagC}, regMap{"A": 0x11, "F": FlagH}, 4},
		{opcode{0xA4}, "H", regMap{"A": 0x13, "H": 0x24, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
		{opcode{0xA5}, "L", regMap{"A": 0x13, "L": 0x35, "F": FlagN | FlagC}, regMap{"A": 0x11, "F": FlagH}, 4},
		{opcode{0xA5}, "L", regMap{"A": 0x13, "L": 0x24, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
		{opcode{0xA7}, "A", regMap{"A": 0x13, "F": FlagN | FlagC}, regMap{"A": 0x13, "F": FlagH}, 4},
		{opcode{0xA7}, "A", regMap{"A": 0x00, "F": FlagN | FlagC}, regMap{"A": 0x00, "F": FlagH | FlagZ}, 4},
	}

	for _, test := range(tests) {
		testDescription := testDescription{
			fmt.Sprintf("'AND %s' executes logical AND of A with register %s and stores result in A", test.srcReg, test.srcReg),
			test.opc,
			test.init,
			test.expected,
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}


func TestAndIndirectExecutesLogicalAndWithA(t *testing.T) {
	testDescriptions := []testDescription{
		testDescription {
			"'AND (HL)' executes logical AND of A with memory content (HL) and stores result in A",
			opcode{0xA6},
			regMap{"A": 0x13, "HL": 0x1234, "F": FlagN | FlagC},
			regMap{"A": 0x11, "F": FlagH},
			memMap{0x1234: 0x35},
			memMap{},
			8,
		},
		testDescription {
			"'AND (HL)' sets Zero flag if result is zero",
			opcode{0xA6},
			regMap{"A": 0x13, "HL": 0x1234, "F": FlagN | FlagC},
			regMap{"A": 0x00, "F": FlagH | FlagZ},
			memMap{0x1234: 0x24},
			memMap{},
			8,
		},
	}

	for _, testDescription := range(testDescriptions) {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}