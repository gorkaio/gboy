package cpu_test

import (
	"testing"
)

func TestJumpRelative(t *testing.T) {
	tests := []testDescription{
		testDescription {
			description: "'JR r8' jumps to a relative positive address from current PC",
			opcode: opcode{0x18, 0x34},
			regsGiven: regMap{"PC": 0x100},
			regsExpected: regMap{"PC": 0x134},
			memReadExpected: memMap{},
			memWriteExpected: memMap{},
			cycles: 12,
		},
		testDescription {
			description: "'JR r8' jumps to a relative negative address from current PC",
			opcode: opcode{0x18, 0xF4},
			regsGiven: regMap{"PC": 0xFF1},
			regsExpected: regMap{"PC": 0xFE5},
			memReadExpected: memMap{},
			memWriteExpected: memMap{},
			cycles: 12,
		},
	}

	for _, test := range(tests) {
		testCase := buildTestCase(test)
		testCase.Run(t)
	}
}