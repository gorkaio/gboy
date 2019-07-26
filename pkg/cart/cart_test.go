package cart

import (
	"gotest.tools/assert"
	"testing"
)

const testfile string = "../../roms/10-print.gb"
const testtitle string = "10 PRINT"

func TestLoadsFile(t *testing.T) {
	cart, err := LoadFromFile(testfile)
	assert.NilError(t, err)
	assert.Equal(t, cart.filename, testfile)
}

func TestReadsCartTitle(t *testing.T) {
	cart, err := LoadFromFile(testfile)
	assert.NilError(t, err)
	assert.Equal(t, cart.Title(), testtitle)
}

func TestReadsControllerType(t *testing.T) {
	cart, err := LoadFromFile(testfile)
	assert.NilError(t, err)
	assert.Equal(t, cart.Type(), "ROM only (0x00)")
}

func TestReadsMemory(t *testing.T) {
	cart, err := LoadFromFile(testfile)
	assert.NilError(t, err)
	assert.Equal(t, cart.Read(0x137), byte('P'))
}
