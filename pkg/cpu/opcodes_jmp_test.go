package cpu_test

import (
	"testing"
)

func TestJumpRelative(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JR r8' jumps to a relative positive address from current PC",
			opcode:           opcode{0x18, 0x34},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x134},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR r8' jumps to a relative negative address from current PC",
			opcode:           opcode{0x18, 0xFC},
			regsGiven:        regMap{"PC": 0x216},
			regsExpected:     regMap{"PC": 0x214},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR 0' loops",
			opcode:           opcode{0x18, 0x00},
			regsGiven:        regMap{"PC": 0x216},
			regsExpected:     regMap{"PC": 0x216},
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

func TestJumpRelativeIfNonZero(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JR NZ, r8' jumps to a relative positive address from current PC if Flag Z not set",
			opcode:           opcode{0x20, 0x34},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x134},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR NZ, r8' jumps to a relative negative address from current PC if Flag Z not set",
			opcode:           opcode{0x20, 0xFC},
			regsGiven:        regMap{"PC": 0x216},
			regsExpected:     regMap{"PC": 0x214},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR NZ, r8' does not jump if Flag Z set",
			opcode:           opcode{0x20, 0x34},
			regsGiven:        regMap{"PC": 0x100, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x102, "F": FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR NZ, r8' does not jump if Flag Z set",
			opcode:           opcode{0x20, 0xF4},
			regsGiven:        regMap{"PC": 0xFF1, "F": FlagZ},
			regsExpected:     regMap{"PC": 0xFF3, "F": FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR NZ, 0' loops if Flag Z clear",
			opcode:           opcode{0x20, 0x00},
			regsGiven:        regMap{"PC": 0xFF1},
			regsExpected:     regMap{"PC": 0xFF1},
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

func TestJumpRelativeIfZero(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JR NZ, r8' jumps to a relative positive address from current PC if Flag Z set",
			opcode:           opcode{0x28, 0x34},
			regsGiven:        regMap{"PC": 0x100, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x134, "F": FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR NZ, r8' jumps to a relative negative address from current PC if Flag Z set",
			opcode:           opcode{0x28, 0xFC},
			regsGiven:        regMap{"PC": 0x216, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x214, "F": FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR NZ, r8' does not jump if Flag Z not set",
			opcode:           opcode{0x28, 0x34},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x102},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR NZ, r8' does not jump if Flag Z not set",
			opcode:           opcode{0x28, 0xF4},
			regsGiven:        regMap{"PC": 0xFF1},
			regsExpected:     regMap{"PC": 0xFF3},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR NZ, 0' loops if Flag Z set",
			opcode:           opcode{0x28, 0x00},
			regsGiven:        regMap{"PC": 0x216, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x216, "F": FlagZ},
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

func TestJumpRelativeIfNonCarry(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JR NC, r8' jumps to a relative positive address from current PC if Flag C not set",
			opcode:           opcode{0x30, 0x34},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x134},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR NC, r8' jumps to a relative negative address from current PC if Flag C not set",
			opcode:           opcode{0x30, 0xFC},
			regsGiven:        regMap{"PC": 0x216},
			regsExpected:     regMap{"PC": 0x214},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR NC, r8' does not jump if Flag C set",
			opcode:           opcode{0x30, 0x34},
			regsGiven:        regMap{"PC": 0x100, "F": FlagC},
			regsExpected:     regMap{"PC": 0x102, "F": FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR NC, r8' does not jump if Flag C set",
			opcode:           opcode{0x30, 0xF4},
			regsGiven:        regMap{"PC": 0xFF1, "F": FlagC},
			regsExpected:     regMap{"PC": 0xFF3, "F": FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR NC, 0' loops if Flag C not set",
			opcode:           opcode{0x30, 0x00},
			regsGiven:        regMap{"PC": 0x216},
			regsExpected:     regMap{"PC": 0x216},
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

func TestJumpRelativeIfCarry(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JR C, r8' jumps to a relative positive address from current PC if Flag C not set",
			opcode:           opcode{0x38, 0x34},
			regsGiven:        regMap{"PC": 0x100, "F": FlagC},
			regsExpected:     regMap{"PC": 0x134, "F": FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR C, r8' jumps to a relative negative address from current PC if Flag C not set",
			opcode:           opcode{0x38, 0xFC},
			regsGiven:        regMap{"PC": 0x216, "F": FlagC},
			regsExpected:     regMap{"PC": 0x214, "F": FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           12,
		},
		{
			description:      "'JR C, r8' does not jump if Flag C set",
			opcode:           opcode{0x38, 0x34},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x102},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR C, r8' does not jump if Flag C set",
			opcode:           opcode{0x38, 0xF4},
			regsGiven:        regMap{"PC": 0xFF1},
			regsExpected:     regMap{"PC": 0xFF3},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'JR C, 0' loops if Flag C not set",
			opcode:           opcode{0x38, 0x00},
			regsGiven:        regMap{"PC": 0x216, "F": FlagC},
			regsExpected:     regMap{"PC": 0x216, "F": FlagC},
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

func TestJumpAbsoluteImmediate(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JP a16' jumps to an absolute address",
			opcode:           opcode{0xC3, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x1234},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           16,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestJumpAbsoluteIndirect(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JP (HL)' jumps to an absolute address contained in register HL",
			opcode:           opcode{0xE9},
			regsGiven:        regMap{"PC": 0x100, "HL": 0x1234},
			regsExpected:     regMap{"PC": 0x1234},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           4,
		},
	}

	for _, test := range tests {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}

func TestJumpAbsoluteIfNonZero(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JP NZ, a16' jumps to an absolute address if Flag Z not set",
			opcode:           opcode{0xC2, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x1234},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           16,
		},
		{
			description:      "'JP NZ, a16' does not jump if Flag Z set",
			opcode:           opcode{0xC2, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x103, "F": FlagZ},
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

func TestJumpAbsoluteIfZero(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JP Z, a16' jumps to an absolute address if Flag Z set",
			opcode:           opcode{0xCA, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100, "F": FlagZ},
			regsExpected:     regMap{"PC": 0x1234, "F": FlagZ},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           16,
		},
		{
			description:      "'JP Z, a16' does not jump if Flag Z not set",
			opcode:           opcode{0xCA, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x103},
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

func TestJumpAbsoluteIfNonCarry(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JP NC, a16' jumps to an absolute address if Flag C not set",
			opcode:           opcode{0xD2, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x1234},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           16,
		},
		{
			description:      "'JP NC, a16' does not jump if Flag C set",
			opcode:           opcode{0xD2, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100, "F": FlagC},
			regsExpected:     regMap{"PC": 0x103, "F": FlagC},
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

func TestJumpAbsoluteIfCarry(t *testing.T) {
	tests := []testDescription{
		{
			description:      "'JP C, a16' jumps to an absolute address if Flag C set",
			opcode:           opcode{0xDA, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100, "F": FlagC},
			regsExpected:     regMap{"PC": 0x1234, "F": FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           16,
		},
		{
			description:      "'JP C, a16' does not jump if Flag C not set",
			opcode:           opcode{0xDA, 0x34, 0x12},
			regsGiven:        regMap{"PC": 0x100},
			regsExpected:     regMap{"PC": 0x103},
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
