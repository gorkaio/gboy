package cpu_test

import (
	"fmt"
	"testing"
)

func TestAddClearsNegativeFlagFor16bitRegisters(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint16
		expectedVal    uint16
		cycles         int
	}{
		{opc: opcode{0x09}, dst: "HL", src: "BC", dstVal: 0x1234, srcVal: 0x1234, expectedVal: 0x2468, cycles: 8},
		{opc: opcode{0x19}, dst: "HL", src: "DE", dstVal: 0x1234, srcVal: 0x1234, expectedVal: 0x2468, cycles: 8},
		{opc: opcode{0x29}, dst: "HL", src: "HL", dstVal: 0x1234, srcVal: 0x1234, expectedVal: 0x2468, cycles: 8},
		{opc: opcode{0x39}, dst: "HL", src: "SP", dstVal: 0x1234, srcVal: 0x1234, expectedVal: 0x2468, cycles: 8},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf(
				"'ADD %s, %s' adds %s to %s and stores result in %s. Negative flag cleared.",
				test.dst,
				test.src,
				test.src,
				test.dst,
				test.dst,
			),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal), "F": FlagN},
			regMap{test.dst: int(test.expectedVal), "F": 0},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddSetsHalfCarryFlagFor16bitRegisters(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint16
		expectedVal    uint16
		cycles         int
	}{
		{opc: opcode{0x09}, dst: "HL", src: "BC", dstVal: 0x0FFF, srcVal: 0x0FFF, expectedVal: 0x1FFE, cycles: 8},
		{opc: opcode{0x19}, dst: "HL", src: "DE", dstVal: 0x0FFF, srcVal: 0x0FFF, expectedVal: 0x1FFE, cycles: 8},
		{opc: opcode{0x29}, dst: "HL", src: "HL", dstVal: 0x0FFF, srcVal: 0x0FFF, expectedVal: 0x1FFE, cycles: 8},
		{opc: opcode{0x39}, dst: "HL", src: "SP", dstVal: 0x0FFF, srcVal: 0x0FFF, expectedVal: 0x1FFE, cycles: 8},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'ADD %s, %s' sets half-carry flag when carry happens in bits 11-12", test.dst, test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal)},
			regMap{test.dst: int(test.expectedVal), "F": FlagH},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddSetsCarryFlagFor16bitRegistersWhenOverflown(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint16
		expectedVal    uint16
		cycles         int
	}{
		{opc: opcode{0x09}, dst: "HL", src: "BC", dstVal: 0x8002, srcVal: 0x8002, expectedVal: 0x0004, cycles: 8},
		{opc: opcode{0x19}, dst: "HL", src: "DE", dstVal: 0x8002, srcVal: 0x8002, expectedVal: 0x0004, cycles: 8},
		{opc: opcode{0x29}, dst: "HL", src: "HL", dstVal: 0x8002, srcVal: 0x8002, expectedVal: 0x0004, cycles: 8},
		{opc: opcode{0x39}, dst: "HL", src: "SP", dstVal: 0x8002, srcVal: 0x8002, expectedVal: 0x0004, cycles: 8},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'ADD %s, %s' sets carry flag when overflown", test.dst, test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal)},
			regMap{test.dst: int(test.expectedVal), "F": FlagC},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddSetsZeroAndCarryFlagFor16bitRegistersWhenOverflownAndResultZero(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint16
		expectedVal    uint16
		cycles         int
	}{
		{opc: opcode{0x09}, dst: "HL", src: "BC", dstVal: 0x8000, srcVal: 0x8000, expectedVal: 0, cycles: 8},
		{opc: opcode{0x19}, dst: "HL", src: "DE", dstVal: 0x8000, srcVal: 0x8000, expectedVal: 0, cycles: 8},
		{opc: opcode{0x29}, dst: "HL", src: "HL", dstVal: 0x8000, srcVal: 0x8000, expectedVal: 0, cycles: 8},
		{opc: opcode{0x39}, dst: "HL", src: "SP", dstVal: 0x8000, srcVal: 0x8000, expectedVal: 0, cycles: 8},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'ADD %s, %s' sets zero flag and carry when overflown and result is zero", test.dst, test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal)},
			regMap{test.dst: int(test.expectedVal), "F": FlagZ | FlagC},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddClearsNegativeFlagFor8bitRegisters(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint8
		expectedVal    uint8
		cycles         int
	}{
		{opc: opcode{0x80}, dst: "A", src: "B", dstVal: 0x12, srcVal: 0x34, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x81}, dst: "A", src: "C", dstVal: 0x12, srcVal: 0x34, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x82}, dst: "A", src: "D", dstVal: 0x12, srcVal: 0x34, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x83}, dst: "A", src: "E", dstVal: 0x12, srcVal: 0x34, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x84}, dst: "A", src: "H", dstVal: 0x12, srcVal: 0x34, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x85}, dst: "A", src: "L", dstVal: 0x12, srcVal: 0x34, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x87}, dst: "A", src: "A", dstVal: 0x12, srcVal: 0x12, expectedVal: 0x24, cycles: 4},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf(
				"'ADD %s, %s' adds %s to %s and stores result in %s. Negative flag cleared.",
				test.dst,
				test.src,
				test.src,
				test.dst,
				test.dst,
			),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal), "F": FlagN},
			regMap{test.dst: int(test.expectedVal), "F": 0},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddSetsHalfCarryFlagFor8bitRegisters(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint8
		expectedVal    uint8
		cycles         int
	}{
		{opc: opcode{0x80}, dst: "A", src: "B", dstVal: 0x0F, srcVal: 0x02, expectedVal: 0x11, cycles: 4},
		{opc: opcode{0x81}, dst: "A", src: "C", dstVal: 0x0F, srcVal: 0x02, expectedVal: 0x11, cycles: 4},
		{opc: opcode{0x82}, dst: "A", src: "D", dstVal: 0x0F, srcVal: 0x02, expectedVal: 0x11, cycles: 4},
		{opc: opcode{0x83}, dst: "A", src: "E", dstVal: 0x0F, srcVal: 0x02, expectedVal: 0x11, cycles: 4},
		{opc: opcode{0x84}, dst: "A", src: "H", dstVal: 0x0F, srcVal: 0x02, expectedVal: 0x11, cycles: 4},
		{opc: opcode{0x85}, dst: "A", src: "L", dstVal: 0x0F, srcVal: 0x02, expectedVal: 0x11, cycles: 4},
		{opc: opcode{0x87}, dst: "A", src: "A", dstVal: 0x0F, srcVal: 0x0F, expectedVal: 0x1E, cycles: 4},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'ADD %s, %s' sets half-carry flag when carry happens in bits 3-4", test.dst, test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal)},
			regMap{test.dst: int(test.expectedVal), "F": FlagH},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddSetsCarryFlagFor8bitRegistersWhenOverflown(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint8
		expectedVal    uint8
		cycles         int
	}{
		{opc: opcode{0x80}, dst: "A", src: "B", dstVal: 0x82, srcVal: 0x83, expectedVal: 0x05, cycles: 4},
		{opc: opcode{0x81}, dst: "A", src: "C", dstVal: 0x82, srcVal: 0x83, expectedVal: 0x05, cycles: 4},
		{opc: opcode{0x82}, dst: "A", src: "D", dstVal: 0x82, srcVal: 0x83, expectedVal: 0x05, cycles: 4},
		{opc: opcode{0x83}, dst: "A", src: "E", dstVal: 0x82, srcVal: 0x83, expectedVal: 0x05, cycles: 4},
		{opc: opcode{0x84}, dst: "A", src: "H", dstVal: 0x82, srcVal: 0x83, expectedVal: 0x05, cycles: 4},
		{opc: opcode{0x85}, dst: "A", src: "L", dstVal: 0x82, srcVal: 0x83, expectedVal: 0x05, cycles: 4},
		{opc: opcode{0x87}, dst: "A", src: "A", dstVal: 0x82, srcVal: 0x82, expectedVal: 0x04, cycles: 4},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'ADD %s, %s' sets carry flag when overflown", test.dst, test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal)},
			regMap{test.dst: int(test.expectedVal), "F": FlagC},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddSetsZeroAndCarryFlagFor8bitRegistersWhenOverflownAndResultZero(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint8
		expectedVal    uint8
		cycles         int
	}{
		{opc: opcode{0x80}, dst: "A", src: "B", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
		{opc: opcode{0x81}, dst: "A", src: "C", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
		{opc: opcode{0x82}, dst: "A", src: "D", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
		{opc: opcode{0x83}, dst: "A", src: "E", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
		{opc: opcode{0x84}, dst: "A", src: "H", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
		{opc: opcode{0x85}, dst: "A", src: "L", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
		{opc: opcode{0x87}, dst: "A", src: "A", dstVal: 0x80, srcVal: 0x80, expectedVal: 0x00, cycles: 4},
	}

	for _, test := range tests {
		testFlagN := testDescription{
			fmt.Sprintf("'ADD %s, %s' sets zero flag and carry when overflown and result is zero", test.dst, test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal)},
			regMap{test.dst: int(test.expectedVal), "F": FlagZ | FlagC},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddIndirectFromMemoryClearsNegativeFlag(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD A, (HL)' adds content from memory address HL to A and stores result in A. Negative flag cleared.",
		opcode:           opcode{0x86},
		regsGiven:        regMap{"A": 0x12, "HL": 0x1234, "F": FlagN},
		regsExpected:     regMap{"A": 0x68, "F": 0},
		memReadExpected:  memMap{0x1234: 0x56},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddIndirectFromMemorySetsHalfCarryFlag(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD A, (HL)' sets half-carry flag when carry in but 3-4.",
		opcode:           opcode{0x86},
		regsGiven:        regMap{"A": 0x0F, "HL": 0x1234},
		regsExpected:     regMap{"A": 0x11, "F": FlagH},
		memReadExpected:  memMap{0x1234: 0x02},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddIndirectFromMemorySetsCarryFlagWhenOverflown(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD A, (HL)' sets carry flag when overflown.",
		opcode:           opcode{0x86},
		regsGiven:        regMap{"A": 0x82, "HL": 0x1234},
		regsExpected:     regMap{"A": 0x05, "F": FlagC},
		memReadExpected:  memMap{0x1234: 0x83},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddIndirectFromMemorySetsCarryAndZeroFlagsWhenOverflownAndResultZero(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD A, (HL)' sets carry and zero flags when overflown and result zero.",
		opcode:           opcode{0x86},
		regsGiven:        regMap{"A": 0x80, "HL": 0x1234},
		regsExpected:     regMap{"A": 0x00, "F": FlagZ | FlagC},
		memReadExpected:  memMap{0x1234: 0x80},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddSignedRelativeToSPClearsNegativeFlag(t *testing.T) {
	testDescription := testDescription{
		description:      fmt.Sprintf("'ADD SP, %#02x' adds signed %#02x value to SP and stores result in SP. Negative flag cleared.", 0x34, 0x34),
		opcode:           opcode{0xE8, 0x34},
		regsGiven:        regMap{"SP": 0x0012, "F": FlagN},
		regsExpected:     regMap{"SP": 0x0046, "F": 0},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           16,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddSignedRelativeToSPDecrementsSPWithNegativeNumbers(t *testing.T) {
	testDescription := testDescription{
		description:      fmt.Sprintf("'ADD SP, %#02x' decrements SP with signed negative numbers (0xFE = -2).", 0xFE),
		opcode:           opcode{0xE8, 0xFC},
		regsGiven:        regMap{"SP": 0x0012},
		regsExpected:     regMap{"SP": 0x0010},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           16,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddSignedRelativeToSPClearsZeroFlagEvenWithResultZero(t *testing.T) {
	testDescription := testDescription{
		description:      fmt.Sprintf("'ADD SP, %#02x' clears zero flag even with result zero.", 0x08),
		opcode:           opcode{0xE8, 0x08},
		regsGiven:        regMap{"SP": 0xFFF8, "F": FlagZ},
		regsExpected:     regMap{"SP": 0x0000, "F": FlagC | FlagH},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           16,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddSignedRelativeToSPSetsHalfCarryFlag(t *testing.T) {
	testDescription := testDescription{
		description:      fmt.Sprintf("'ADD SP, %#02x' sets half-carry flag when carry in bits 3-4.", 0x02),
		opcode:           opcode{0xE8, 0x02},
		regsGiven:        regMap{"SP": 0x0FFF},
		regsExpected:     regMap{"SP": 0x1001, "F": FlagH},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           16,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddSignedRelativeToSPSetsCarryFlagWhenOverflown(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD SP, %#02x' sets carry flag when overflown.",
		opcode:           opcode{0xE8, 0x03},
		regsGiven:        regMap{"SP": 0xFFFE},
		regsExpected:     regMap{"SP": 0x0001, "F": FlagC | FlagH},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           16,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddImmediateClearsNegativeFlag(t *testing.T) {
	testDescription := testDescription{
		description:      fmt.Sprintf("'ADD A, %#02x' adds %#02x value to A and stores result in A. Negative flag cleared.", 0x34, 0x34),
		opcode:           opcode{0xC6, 0x34},
		regsGiven:        regMap{"A": 0x12, "F": FlagN},
		regsExpected:     regMap{"A": 0x46, "F": 0},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddImmediateSetsHalfCarryFlag(t *testing.T) {
	testDescription := testDescription{
		description:      fmt.Sprintf("'ADD A, %#02x' sets half-carry flag when carry in bits 3-4.", 0x02),
		opcode:           opcode{0xC6, 0x02},
		regsGiven:        regMap{"A": 0x0F},
		regsExpected:     regMap{"A": 0x11, "F": FlagH},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddImmediateSetsCarryFlagWhenOverflown(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD A, (HL)' sets carry flag when overflown.",
		opcode:           opcode{0xC6, 0x83},
		regsGiven:        regMap{"A": 0x82},
		regsExpected:     regMap{"A": 0x05, "F": FlagC},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddImmediateSetsCarryAndZeroFlagsWhenOverflownAndResultZero(t *testing.T) {
	testDescription := testDescription{
		description:      "'ADD A, (HL)' sets carry and zero flags when overflown and result zero.",
		opcode:           opcode{0xC6, 0x80},
		regsGiven:        regMap{"A": 0x80},
		regsExpected:     regMap{"A": 0x00, "F": FlagZ | FlagC},
		memReadExpected:  memMap{},
		memWriteExpected: memMap{},
		cycles:           8,
	}
	testCase := buildTestCase(testDescription)
	testCase.Run(t)
}

func TestAddImmediateWithCarryClearsNegativeFlag(t *testing.T) {
	tests := []testDescription{
		{
			description:      fmt.Sprintf("'ADC %#02x' adds %#02x value to A and stores result in A. Negative flag cleared.", 0x34, 0x34),
			opcode:           opcode{0xCE, 0x34},
			regsGiven:        regMap{"A": 0x12, "F": FlagN},
			regsExpected:     regMap{"A": 0x46, "F": 0},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      fmt.Sprintf("'ADC %#02x' adds %#02x value to A with carry and stores result in A. Negative flag cleared.", 0x34, 0x34),
			opcode:           opcode{0xCE, 0x34},
			regsGiven:        regMap{"A": 0x12, "F": FlagN | FlagC},
			regsExpected:     regMap{"A": 0x47, "F": 0},
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

func TestAddImmediateWithCarrySetsHalfCarryFlag(t *testing.T) {
	tests := []testDescription{
		{
			description:      fmt.Sprintf("'ADC %#02x' sets half-carry flag when carry in bits 3-4.", 0x02),
			opcode:           opcode{0xCE, 0x02},
			regsGiven:        regMap{"A": 0x0F},
			regsExpected:     regMap{"A": 0x11, "F": FlagH},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      fmt.Sprintf("'ADC %#02x' sets half-carry flag with carry when carry in bits 3-4.", 0x02),
			opcode:           opcode{0xCE, 0x02},
			regsGiven:        regMap{"A": 0x0F, "F": FlagC},
			regsExpected:     regMap{"A": 0x12, "F": FlagH},
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

func TestAddImmediateWithCarrySetsCarryFlagWhenOverflown(t *testing.T) {
	tests := []testDescription{
		{
			description:      fmt.Sprintf("'ADC %#02x' sets carry flag when overflown.", 0x83),
			opcode:           opcode{0xCE, 0x83},
			regsGiven:        regMap{"A": 0x82},
			regsExpected:     regMap{"A": 0x05, "F": FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      fmt.Sprintf("'ADC %#02x' sets carry flag with carry when overflown.", 0x33),
			opcode:           opcode{0xCE, 0x83},
			regsGiven:        regMap{"A": 0x82, "F": FlagC},
			regsExpected:     regMap{"A": 0x06, "F": FlagC},
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

func TestAddImmediateWithCarrySetsCarryAndZeroFlagsWhenOverflownAndResultZero(t *testing.T) {
	tests := []testDescription{
		{
			description:      fmt.Sprintf("'ADC %#02x' sets carry and zero flags when overflown and result zero.", 0x80),
			opcode:           opcode{0xCE, 0x80},
			regsGiven:        regMap{"A": 0x80},
			regsExpected:     regMap{"A": 0x00, "F": FlagZ | FlagC},
			memReadExpected:  memMap{},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      fmt.Sprintf("'ADC %#02x' sets carry and zero flags with carry when overflown and result zero.", 0x80),
			opcode:           opcode{0xCE, 0x80},
			regsGiven:        regMap{"A": 0x7F, "F": FlagC},
			regsExpected:     regMap{"A": 0x00, "F": FlagZ | FlagC | FlagH},
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

func TestAddWithCarryDirectAddsCarryWhenSet(t *testing.T) {
	tests := []struct {
		opc            opcode
		src, dst       string
		srcVal, dstVal uint8
		carry          bool
		expectedVal    uint8
		cycles         int
	}{
		{opc: opcode{0x88}, dst: "A", src: "B", dstVal: 0x12, srcVal: 0x34, carry: true, expectedVal: 0x47, cycles: 4},
		{opc: opcode{0x89}, dst: "A", src: "C", dstVal: 0x12, srcVal: 0x34, carry: true, expectedVal: 0x47, cycles: 4},
		{opc: opcode{0x8A}, dst: "A", src: "D", dstVal: 0x12, srcVal: 0x34, carry: true, expectedVal: 0x47, cycles: 4},
		{opc: opcode{0x8B}, dst: "A", src: "E", dstVal: 0x12, srcVal: 0x34, carry: true, expectedVal: 0x47, cycles: 4},
		{opc: opcode{0x8C}, dst: "A", src: "H", dstVal: 0x12, srcVal: 0x34, carry: true, expectedVal: 0x47, cycles: 4},
		{opc: opcode{0x8D}, dst: "A", src: "L", dstVal: 0x12, srcVal: 0x34, carry: true, expectedVal: 0x47, cycles: 4},
		{opc: opcode{0x8F}, dst: "A", src: "A", dstVal: 0x12, srcVal: 0x12, carry: true, expectedVal: 0x25, cycles: 4},
		{opc: opcode{0x88}, dst: "A", src: "B", dstVal: 0x12, srcVal: 0x34, carry: false, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x89}, dst: "A", src: "C", dstVal: 0x12, srcVal: 0x34, carry: false, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x8A}, dst: "A", src: "D", dstVal: 0x12, srcVal: 0x34, carry: false, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x8B}, dst: "A", src: "E", dstVal: 0x12, srcVal: 0x34, carry: false, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x8C}, dst: "A", src: "H", dstVal: 0x12, srcVal: 0x34, carry: false, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x8D}, dst: "A", src: "L", dstVal: 0x12, srcVal: 0x34, carry: false, expectedVal: 0x46, cycles: 4},
		{opc: opcode{0x8F}, dst: "A", src: "A", dstVal: 0x12, srcVal: 0x12, carry: false, expectedVal: 0x24, cycles: 4},
	}

	for _, test := range tests {
		var regF uint8 = FlagN
		if test.carry {
			regF = regF | FlagC
		}

		testFlagN := testDescription{
			fmt.Sprintf("'ADC A, %s' adds with carry when flag C is set", test.src),
			test.opc,
			regMap{test.src: int(test.srcVal), test.dst: int(test.dstVal), "F": int(regF)},
			regMap{test.dst: int(test.expectedVal), "F": 0},
			memMap{},
			memMap{},
			test.cycles,
		}
		testCase := buildTestCase(testFlagN)
		testCase.Run(t)
	}
}

func TestAddWithCarryIndirectAddsCarryWhenSet(t *testing.T) {
	testDescriptions := []testDescription{
		{
			description:      "'ADC A, (HL)' adds with carry when flag C is set",
			opcode:           opcode{0x8E},
			regsGiven:        regMap{"A": 0x04, "HL": 0x1234, "F": FlagN},
			regsExpected:     regMap{"A": 0x5A, "F": 0},
			memReadExpected:  memMap{0x1234: 0x56},
			memWriteExpected: memMap{},
			cycles:           8,
		},
		{
			description:      "'ADC A, (HL)' adds with carry when flag C is set",
			opcode:           opcode{0x8E},
			regsGiven:        regMap{"A": 0x04, "HL": 0x1234, "F": FlagN | FlagC},
			regsExpected:     regMap{"A": 0x5B, "F": 0},
			memReadExpected:  memMap{0x1234: 0x56},
			memWriteExpected: memMap{},
			cycles:           8,
		},
	}

	for _, testDescription := range testDescriptions {
		testCase := buildTestCase(testDescription)
		testCase.Run(t)
	}
}
