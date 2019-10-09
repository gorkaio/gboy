package cpu_test

import (
	"testing"
)

func TestNopDoesNotChangeRegistersOrMemory(t *testing.T) {
	testDescription := testDescription{
		"NOP does not interact with registers or memory",
		opcode{0x00},
		regMap{},
		regMap{},
		memMap{},
		memMap{},
		4,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}