package cpu_test

import (
	"fmt"
	"testing"
)

func TestDirect8bitRegisterSubstraction(t *testing.T) {
	tests := []struct {
		opc          opcode
		srcReg       string
		init, expect regMap
		cycles       int
	}{
		{opcode{0x90}, "B", regMap{"A": 0xCF, "B": 0x02}, regMap{"A": 0xCD, "F": FlagN}, 4},
		{opcode{0x90}, "B", regMap{"A": 0x02, "B": 0xCF}, regMap{"A": 0x33, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x90}, "B", regMap{"A": 0x02, "B": 0x02}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x91}, "C", regMap{"A": 0xCF, "C": 0x02}, regMap{"A": 0xCD, "F": FlagN}, 4},
		{opcode{0x91}, "C", regMap{"A": 0x02, "C": 0xCF}, regMap{"A": 0x33, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x91}, "C", regMap{"A": 0x02, "C": 0x02}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x92}, "D", regMap{"A": 0xCF, "D": 0x02}, regMap{"A": 0xCD, "F": FlagN}, 4},
		{opcode{0x92}, "D", regMap{"A": 0x02, "D": 0xCF}, regMap{"A": 0x33, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x92}, "D", regMap{"A": 0x02, "D": 0x02}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x93}, "E", regMap{"A": 0xCF, "E": 0x02}, regMap{"A": 0xCD, "F": FlagN}, 4},
		{opcode{0x93}, "E", regMap{"A": 0x02, "E": 0xCF}, regMap{"A": 0x33, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x93}, "E", regMap{"A": 0x02, "E": 0x02}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x94}, "H", regMap{"A": 0xCF, "H": 0x02}, regMap{"A": 0xCD, "F": FlagN}, 4},
		{opcode{0x94}, "H", regMap{"A": 0x02, "H": 0xCF}, regMap{"A": 0x33, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x94}, "H", regMap{"A": 0x02, "H": 0x02}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x95}, "L", regMap{"A": 0xCF, "L": 0x02}, regMap{"A": 0xCD, "F": FlagN}, 4},
		{opcode{0x95}, "L", regMap{"A": 0x02, "L": 0xCF}, regMap{"A": 0x33, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x95}, "L", regMap{"A": 0x02, "L": 0x02}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x97}, "A", regMap{"A": 0xCF}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'SUB %s' substracts %s from A and stores result in A. Sets negative flag.", test.srcReg, test.srcReg),
			test.opc,
			test.init,
			test.expect,
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestDirect8bitRegisterSubstractionWithCarry(t *testing.T) {
	tests := []struct {
		opc          opcode
		srcReg       string
		init, expect regMap
		cycles       int
	}{
		{opcode{0x98}, "B", regMap{"A": 0xCF, "B": 0x02, "F": FlagC}, regMap{"A": 0xCC, "F": FlagN}, 4},
		{opcode{0x98}, "B", regMap{"A": 0x0A, "B": 0xCE, "F": FlagC}, regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x98}, "B", regMap{"A": 0x02, "B": 0x01, "F": FlagC}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x99}, "C", regMap{"A": 0xCF, "C": 0x02, "F": FlagC}, regMap{"A": 0xCC, "F": FlagN}, 4},
		{opcode{0x99}, "C", regMap{"A": 0x0A, "C": 0xCE, "F": FlagC}, regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x99}, "C", regMap{"A": 0x02, "C": 0x01, "F": FlagC}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x9A}, "D", regMap{"A": 0xCF, "D": 0x02, "F": FlagC}, regMap{"A": 0xCC, "F": FlagN}, 4},
		{opcode{0x9A}, "D", regMap{"A": 0x0A, "D": 0xCE, "F": FlagC}, regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x9A}, "D", regMap{"A": 0x02, "D": 0x01, "F": FlagC}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x9B}, "E", regMap{"A": 0xCF, "E": 0x02, "F": FlagC}, regMap{"A": 0xCC, "F": FlagN}, 4},
		{opcode{0x9B}, "E", regMap{"A": 0x0A, "E": 0xCE, "F": FlagC}, regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x9B}, "E", regMap{"A": 0x02, "E": 0x01, "F": FlagC}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x9C}, "H", regMap{"A": 0xCF, "H": 0x02, "F": FlagC}, regMap{"A": 0xCC, "F": FlagN}, 4},
		{opcode{0x9C}, "H", regMap{"A": 0x0A, "H": 0xCE, "F": FlagC}, regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x9C}, "H", regMap{"A": 0x02, "H": 0x01, "F": FlagC}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x9D}, "H", regMap{"A": 0xCF, "L": 0x02, "F": FlagC}, regMap{"A": 0xCC, "F": FlagN}, 4},
		{opcode{0x9D}, "H", regMap{"A": 0x0A, "L": 0xCE, "F": FlagC}, regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH}, 4},
		{opcode{0x9D}, "H", regMap{"A": 0x02, "L": 0x01, "F": FlagC}, regMap{"A": 0x00, "F": FlagN | FlagZ}, 4},
		{opcode{0x9F}, "H", regMap{"A": 0xCF, "F": FlagC}, regMap{"A": 0xFF, "F": FlagN | FlagC}, 4},
		{opcode{0x9F}, "H", regMap{"A": 0x02, "F": FlagC}, regMap{"A": 0xFF, "F": FlagN | FlagC | FlagH}, 4},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'SBC A, %s' substracts %s from A and stores result in A. Sets negative flag.", test.srcReg, test.srcReg),
			test.opc,
			test.init,
			test.expect,
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestIndirect8bitRegisterSubstraction(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'SUB (HL)' substracts content from memory address HL form A and stores result in A. Sets negative flag.",
			opcode:           opcode{0x96},
			regsGiven:        regMap{"A": 0xCF, "HL": 0x1234},
			regsExpected:     regMap{"A": 0xCD, "F": FlagN},
			memReadExpected:  memMap{0x1234: 0x02},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SUB (HL)' sets arithmetic flags.",
			opcode:           opcode{0x96},
			regsGiven:        regMap{"A": 0x0A, "HL": 0x1234},
			regsExpected:     regMap{"A": 0x3C, "F": FlagN | FlagC | FlagH},
			memReadExpected:  memMap{0x1234: 0xCE},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SUB (HL)' sets zero flag if result is zero.",
			opcode:           opcode{0x96},
			regsGiven:        regMap{"A": 0x01, "HL": 0x1234},
			regsExpected:     regMap{"A": 0x00, "F": FlagN | FlagZ},
			memReadExpected:  memMap{0x1234: 0x01},
			memWriteExpected: memMap{},
			cycles:           8,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestIndirect8bitRegisterSubstractionWithCarry(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'SBC A, (HL)' substracts content from memory address HL form A and stores result in A without carry. Sets negative flag.",
			opcode:           opcode{0x9E},
			regsGiven:        regMap{"A": 0xCF, "HL": 0x1234},
			regsExpected:     regMap{"A": 0xCD, "F": FlagN},
			memReadExpected:  memMap{0x1234: 0x02},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, (HL)' substracts content from memory address HL form A and stores result in A with carry. Sets negative flag.",
			opcode:           opcode{0x9E},
			regsGiven:        regMap{"A": 0xCF, "HL": 0x1234, "F": FlagC},
			regsExpected:     regMap{"A": 0xCC, "F": FlagN},
			memReadExpected:  memMap{0x1234: 0x02},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, (HL)' sets arithmetic flags without carry.",
			opcode:           opcode{0x9E},
			regsGiven:        regMap{"A": 0x02, "HL": 0x1234},
			regsExpected:     regMap{"A": 0x33, "F": FlagN | FlagC | FlagH},
			memReadExpected:  memMap{0x1234: 0xCF},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, (HL)' sets arithmetic flags with carry.",
			opcode:           opcode{0x9E},
			regsGiven:        regMap{"A": 0x02, "HL": 0x1234, "F": FlagC},
			regsExpected:     regMap{"A": 0x32, "F": FlagN | FlagC},
			memReadExpected:  memMap{0x1234: 0xCF},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, (HL)' sets zero flag if result is zero without carry.",
			opcode:           opcode{0x9E},
			regsGiven:        regMap{"A": 0x02, "HL": 0x1234},
			regsExpected:     regMap{"A": 0x00, "F": FlagN | FlagZ},
			memReadExpected:  memMap{0x1234: 0x02},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, (HL)' sets zero flag if result is zero with carry.",
			opcode:           opcode{0x9E},
			regsGiven:        regMap{"A": 0x02, "HL": 0x1234, "F": FlagC},
			regsExpected:     regMap{"A": 0x00, "F": FlagN | FlagZ},
			memReadExpected:  memMap{0x1234: 0x01},
			memWriteExpected: memMap{},
			cycles:           8,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestImmediate8bitRegisterSubstraction(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'SBC A, d8' substracts d8 from A and stores result in A. Sets negative flag.",
			opcode:           opcode{0xD6, 0x02},
			regsGiven:        regMap{"A": 0xCF},
			regsExpected:     regMap{"A": 0xCD, "F": FlagN},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, d8' sets arithmetic flags.",
			opcode:           opcode{0xD6, 0xCF},
			regsGiven:        regMap{"A": 0x02},
			regsExpected:     regMap{"A": 0x33, "F": FlagN | FlagC | FlagH},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC A, d8' sets zero flag if result is zero.",
			opcode:           opcode{0xD6, 0x02},
			regsGiven:        regMap{"A": 0x02},
			regsExpected:     regMap{"A": 0x00, "F": FlagN | FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestImmediate8bitRegisterSubstractionWithCarry(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'SBC d8' substracts d8 from A and stores result in A. Sets negative flag.",
			opcode:           opcode{0xDE, 0x02},
			regsGiven:        regMap{"A": 0xCF, "F": FlagC},
			regsExpected:     regMap{"A": 0xCC, "F": FlagN},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC d8' sets arithmetic flags.",
			opcode:           opcode{0xDE, 0xCE},
			regsGiven:        regMap{"A": 0x0A, "HL": 0x1234, "F": FlagC},
			regsExpected:     regMap{"A": 0x3B, "F": FlagN | FlagC | FlagH},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'SBC d8' sets zero flag if result is zero.",
			opcode:           opcode{0xDE, 0x01},
			regsGiven:        regMap{"A": 0x02, "HL": 0x1234, "F": FlagC},
			regsExpected:     regMap{"A": 0x00, "F": FlagN | FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}
