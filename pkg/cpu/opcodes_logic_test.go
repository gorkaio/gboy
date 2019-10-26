package cpu_test

import (
	"fmt"
	"testing"
)

func TestAndDirectExecutesLogicalAndWithA(t *testing.T) {
	tests := []struct {
		opc      opcode
		srcReg   string
		init     regMap
		expected regMap
		cycles   int
	}{
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

	for _, test := range tests {
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
		{
			"'AND (HL)' executes logical AND of A with memory content (HL) and stores result in A",
			opcode{0xA6},
			regMap{"A": 0x13, "HL": 0x1234, "F": FlagN | FlagC},
			regMap{"A": 0x11, "F": FlagH},
			memMap{0x1234: 0x35},
			memMap{},
			8,
		},
		{
			"'AND (HL)' sets Zero flag if result is zero",
			opcode{0xA6},
			regMap{"A": 0x13, "HL": 0x1234, "F": FlagN | FlagC},
			regMap{"A": 0x00, "F": FlagH | FlagZ},
			memMap{0x1234: 0x24},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestOrDirectExecutesLogicalOrWithA(t *testing.T) {
	tests := []struct {
		opc      opcode
		srcReg   string
		init     regMap
		expected regMap
		cycles   int
	}{
		{opcode{0xB0}, "B", regMap{"A": 0x43, "B": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x77, "F": 0}, 4},
		{opcode{0xB0}, "B", regMap{"A": 0x00, "B": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xB1}, "C", regMap{"A": 0x43, "C": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x77, "F": 0}, 4},
		{opcode{0xB1}, "C", regMap{"A": 0x00, "C": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xB2}, "D", regMap{"A": 0x43, "D": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x77, "F": 0}, 4},
		{opcode{0xB2}, "D", regMap{"A": 0x00, "D": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xB3}, "E", regMap{"A": 0x43, "E": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x77, "F": 0}, 4},
		{opcode{0xB3}, "E", regMap{"A": 0x00, "E": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xB4}, "H", regMap{"A": 0x43, "H": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x77, "F": 0}, 4},
		{opcode{0xB4}, "H", regMap{"A": 0x00, "H": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xB5}, "L", regMap{"A": 0x43, "L": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x77, "F": 0}, 4},
		{opcode{0xB5}, "L", regMap{"A": 0x00, "L": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xB7}, "A", regMap{"A": 0x43, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x43, "F": 0}, 4},
		{opcode{0xB7}, "A", regMap{"A": 0x00, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
	}

	for _, test := range tests {
		testDescription := testDescription{
			fmt.Sprintf("'OR %s' executes logical OR of A with register %s and stores result in A", test.srcReg, test.srcReg),
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

func TestOrImmediateExecutesLogicalOrWithA(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'OR d8' executes logical OR of A with d8 and stores result in A",
			opcode{0xF6, 0x35},
			regMap{"A": 0x43, "F": FlagN | FlagC | FlagH},
			regMap{"A": 0x77, "F": 0},
			memMap{},
			memMap{},
			8,
		},
		{
			"'OR d8' sets Zero flag if result is zero",
			opcode{0xF6, 0x00},
			regMap{"A": 0x00, "F": FlagN | FlagC},
			regMap{"A": 0x00, "F": FlagZ},
			memMap{},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestOrIndirectExecutesLogicalOrWithA(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'OR (HL)' executes logical OR of A with memory content (HL) and stores result in A",
			opcode{0xB6},
			regMap{"A": 0x43, "HL": 0x1234, "F": FlagN | FlagC | FlagH},
			regMap{"A": 0x77, "F": 0},
			memMap{0x1234: 0x35},
			memMap{},
			8,
		},
		{
			"'OR (HL)' sets Zero flag if result is zero",
			opcode{0xB6},
			regMap{"A": 0x00, "HL": 0x1234, "F": FlagN | FlagC},
			regMap{"A": 0x00, "F": FlagZ},
			memMap{0x1234: 0x00},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestXorDirectExecutesLogicalXorWithA(t *testing.T) {
	tests := []struct {
		opc      opcode
		srcReg   string
		init     regMap
		expected regMap
		cycles   int
	}{
		{opcode{0xA8}, "B", regMap{"A": 0x43, "B": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x76, "F": 0}, 4},
		{opcode{0xA8}, "B", regMap{"A": 0xBA, "B": 0xBA, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xA9}, "C", regMap{"A": 0x43, "C": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x76, "F": 0}, 4},
		{opcode{0xA9}, "C", regMap{"A": 0xBA, "C": 0xBA, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xAA}, "D", regMap{"A": 0x43, "D": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x76, "F": 0}, 4},
		{opcode{0xAA}, "D", regMap{"A": 0xBA, "D": 0xBA, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xAB}, "E", regMap{"A": 0x43, "E": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x76, "F": 0}, 4},
		{opcode{0xAB}, "E", regMap{"A": 0xBA, "E": 0xBA, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xAC}, "H", regMap{"A": 0x43, "H": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x76, "F": 0}, 4},
		{opcode{0xAC}, "H", regMap{"A": 0xBA, "H": 0xBA, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xAD}, "L", regMap{"A": 0x43, "L": 0x35, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x76, "F": 0}, 4},
		{opcode{0xAD}, "L", regMap{"A": 0xBA, "L": 0xBA, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
		{opcode{0xAF}, "A", regMap{"A": 0x43, "F": FlagN | FlagH | FlagC}, regMap{"A": 0x00, "F": FlagZ}, 4},
	}

	for _, test := range tests {
		testDescription := testDescription{
			fmt.Sprintf("'XOR %s' executes logical XOR of A with register %s and stores result in A", test.srcReg, test.srcReg),
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

func TestXorImmediateExecutesLogicalXorWithA(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'XOR d8' executes logical XOR of A with d8 and stores result in A",
			opcode{0xEE, 0x35},
			regMap{"A": 0x43, "F": FlagN | FlagC | FlagH},
			regMap{"A": 0x76, "F": 0},
			memMap{},
			memMap{},
			8,
		},
		{
			"'XOR d8' sets Zero flag if result is zero",
			opcode{0xEE, 0x56},
			regMap{"A": 0x56, "F": FlagN | FlagC},
			regMap{"A": 0x00, "F": FlagZ},
			memMap{},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestXorIndirectExecutesLogicalXorWithA(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'XOR (HL)' executes logical XOR of A with memory content (HL) and stores result in A",
			opcode{0xAE},
			regMap{"A": 0x43, "HL": 0x1234, "F": FlagN | FlagC | FlagH},
			regMap{"A": 0x76, "F": 0},
			memMap{0x1234: 0x35},
			memMap{},
			8,
		},
		{
			"'XOR (HL)' sets Zero flag if result is zero",
			opcode{0xAE},
			regMap{"A": 0x56, "HL": 0x1234, "F": FlagN | FlagC},
			regMap{"A": 0x00, "F": FlagZ},
			memMap{0x1234: 0x56},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestCpComparesGivenRegisterWithA(t *testing.T) {
	tests := []struct {
		opc      opcode
		srcReg   string
		init     regMap
		expected regMap
		cycles   int
	}{
		{opcode{0xB8}, "B", regMap{"A": 0x43, "B": 0x35}, regMap{"A": 0x43, "F": FlagN | FlagH}, 4},
		{opcode{0xB8}, "B", regMap{"A": 0x43, "B": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
		{opcode{0xB9}, "C", regMap{"A": 0x43, "C": 0x35}, regMap{"A": 0x43, "F": FlagN | FlagH}, 4},
		{opcode{0xB9}, "C", regMap{"A": 0x43, "C": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
		{opcode{0xBA}, "D", regMap{"A": 0x43, "D": 0x35}, regMap{"A": 0x43, "F": FlagN | FlagH}, 4},
		{opcode{0xBA}, "D", regMap{"A": 0x43, "D": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
		{opcode{0xBB}, "E", regMap{"A": 0x43, "E": 0x35}, regMap{"A": 0x43, "F": FlagN | FlagH}, 4},
		{opcode{0xBB}, "E", regMap{"A": 0x43, "E": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
		{opcode{0xBC}, "H", regMap{"A": 0x43, "H": 0x35}, regMap{"A": 0x43, "F": FlagN | FlagH}, 4},
		{opcode{0xBC}, "H", regMap{"A": 0x43, "H": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
		{opcode{0xBD}, "L", regMap{"A": 0x43, "L": 0x35}, regMap{"A": 0x43, "F": FlagN | FlagH}, 4},
		{opcode{0xBD}, "L", regMap{"A": 0x43, "L": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
		{opcode{0xBF}, "A", regMap{"A": 0x43}, regMap{"A": 0x43, "F": FlagN | FlagZ}, 4},
	}

	for _, test := range tests {
		testDescription := testDescription{
			fmt.Sprintf("'CP %s' substracts value from A without storing value and sets flags accordingly", test.srcReg),
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


func TestCpIndirectComparesMemoryValueWithA(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'CP (HL)' substracts value at memory address (HL) from A without storing value and sets flags accordingly",
			opcode{0xBE},
			regMap{"A": 0x43, "HL": 0x1234},
			regMap{"A": 0x43, "F": FlagN},
			memMap{0x1234: 0x11},
			memMap{},
			8,
		},
		{
			"'CP (HL)' substracts value at memory address (HL) from A without storing value and sets flags accordingly",
			opcode{0xBE},
			regMap{"A": 0x56, "HL": 0x1234},
			regMap{"A": 0x56, "F": FlagN | FlagZ},
			memMap{0x1234: 0x56},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}

func TestCpImmediateComparesValueWithA(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"'CP d8' substracts value d8 from A without storing value and sets flags accordingly",
			opcode{0xFE, 0x11},
			regMap{"A": 0x43},
			regMap{"A": 0x43, "F": FlagN},
			memMap{},
			memMap{},
			8,
		},
		{
			"'CP (HL)' substracts value d8 from A without storing value and sets flags accordingly",
			opcode{0xFE, 0x56},
			regMap{"A": 0x56},
			regMap{"A": 0x56, "F": FlagN | FlagZ},
			memMap{},
			memMap{},
			8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}
