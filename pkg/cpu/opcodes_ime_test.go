package cpu_test

import (
	"testing"
)

func TestEIEnablesInterrupts(t *testing.T) {
	testDescription := testDescription{
		"EI enables interrupts",
		opcode{0xFB},
		regMap{},
		regMap{},
		memMap{},
		memMap{},
		4,
	}
	testCase := buildTestCase(testDescription)
	testCase.WithIME(false)
	testCase.ExpectIME(true)
	testCase.Run(t)
}

func TestDIDisablesInterrupts(t *testing.T) {
	testDescription := testDescription{
		"EI enables interrupts",
		opcode{0xF3},
		regMap{},
		regMap{},
		memMap{},
		memMap{},
		4,
	}
	testCase := buildTestCase(testDescription)
	testCase.WithIME(true)
	testCase.ExpectIME(false)
	testCase.Run(t)
}
