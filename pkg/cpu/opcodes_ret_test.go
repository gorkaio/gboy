package cpu_test

import (
	"testing"
)

func TestRetJumpsToSPAddress(t *testing.T) {
	testDescription := testDescription{
		"RET jumps to address contained in memory position (SP) and increments SP",
		opcode{0xC9},
		regMap{"PC": 0x100, "SP": 0x1234},
		regMap{"PC": 0x7856, "SP": 0x1236},
		memMap{0x1234: 0x56, 0x1235: 0x78},
		memMap{},
		16,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestRetiJumpsToSPAddressAndEnablesInterrupts(t *testing.T) {
	testDescription := testDescription{
		"RETI jumps to address contained in memory position (SP), increments SP and enables interrupts",
		opcode{0xD9},
		regMap{"PC": 0x100, "SP": 0x1234},
		regMap{"PC": 0x7856, "SP": 0x1236},
		memMap{0x1234: 0x56, 0x1235: 0x78},
		memMap{},
		16,
	}
	testCase := buildTestCase(testDescription)
	testCase.WithIME(false)
	testCase.ExpectIME(true)
	testCase.Run(t)
}

func TestRetCJumpsToSPAddressIfCarryFlag(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"RETZ jumps to address contained in memory position (SP) and increments SP if Carry Flag",
			opcode{0xD8},
			regMap{"PC": 0x100, "SP": 0x1234, "F": FlagC},
			regMap{"PC": 0x7856, "SP": 0x1236},
			memMap{0x1234: 0x56, 0x1235: 0x78},
			memMap{},
			20,
		},
		{
			"RETZ does not jump if Carry Flag inactive",
			opcode{0xD8},
			regMap{"PC": 0x100, "SP": 0x1234},
			regMap{"PC": 0x101, "SP": 0x1234},
			memMap{},
			memMap{},
			8,
		},
	}
	for _, test := range(testDescriptions) {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestRetNCJumpsToSPAddressIfCarryFlag(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"RETNC jumps to address contained in memory position (SP) and increments SP if not Carry Flag",
			opcode{0xD0},
			regMap{"PC": 0x100, "SP": 0x1234},
			regMap{"PC": 0x7856, "SP": 0x1236},
			memMap{0x1234: 0x56, 0x1235: 0x78},
			memMap{},
			20,
		},
		{
			"RETNC does not jump if Carry Flag set",
			opcode{0xD0},
			regMap{"PC": 0x100, "SP": 0x1234, "F": FlagC},
			regMap{"PC": 0x101, "SP": 0x1234},
			memMap{},
			memMap{},
			8,
		},
	}
	for _, test := range(testDescriptions) {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestRetZJumpsToSPAddressIfZeroFlag(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"RETZ jumps to address contained in memory position (SP) and increments SP if Zero Flag",
			opcode{0xC8},
			regMap{"PC": 0x100, "SP": 0x1234, "F": FlagZ},
			regMap{"PC": 0x7856, "SP": 0x1236},
			memMap{0x1234: 0x56, 0x1235: 0x78},
			memMap{},
			20,
		},
		{
			"RETZ does not jump if Zero Flag inactive",
			opcode{0xC8},
			regMap{"PC": 0x100, "SP": 0x1234},
			regMap{"PC": 0x101, "SP": 0x1234},
			memMap{},
			memMap{},
			8,
		},
	}
	for _, test := range(testDescriptions) {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestRetNZJumpsToSPAddressIfNotZeroFlag(t *testing.T) {
	testDescriptions := []testDescription{
		{
			"RETNZ jumps to address contained in memory position (SP) and increments SP if not Zero Flag",
			opcode{0xC0},
			regMap{"PC": 0x100, "SP": 0x1234},
			regMap{"PC": 0x7856, "SP": 0x1236},
			memMap{0x1234: 0x56, 0x1235: 0x78},
			memMap{},
			20,
		},
		{
			"RETZ does not jump if Zero Flag set",
			opcode{0xC0},
			regMap{"PC": 0x100, "SP": 0x1234, "F": FlagZ},
			regMap{"PC": 0x101, "SP": 0x1234},
			memMap{},
			memMap{},
			8,
		},
	}
	for _, test := range(testDescriptions) {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}