package cpu_test

import (
	"testing"
)

func TestCallNZJumpsAndSetsSPWhenZeroFlagClear(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'CALL NZ, a16' jumps to a16 address with PC in SP if Zero Flag clear",
			opcode:           opcode{0xC4, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100, "SP": 0x402},
			regsExpected:     regMap{"PC": 0x1234, "SP": 0x400},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{0x401: 0x01, 0x400: 0x00},
			cycles:           24,
		},
		{
			description:      "'CALL NZ, a16' does not jump to a16 if Zero Flag set",
			opcode:           opcode{0xC4, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x216, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x219},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}

}

func TestCallNCJumpsAndSetsSPWhenCarryFlagClear(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'CALL NC, a16' jumps to a16 address with PC in SP if Carry Flag clear",
			opcode:           opcode{0xD4, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100, "SP": 0x402},
			regsExpected:     regMap{"PC": 0x1234, "SP": 0x400},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{0x401: 0x01, 0x400: 0x00},
			cycles:           24,
		},
		{
			description:      "'CALL NC, a16' does not jump to a16 if Carry Flag set",
			opcode:           opcode{0xD4, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x216, "F": FlagC},
			regsExpected:     regMap{"PC": 0x219},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}