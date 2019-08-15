package cart_test

import (
	cart "github.com/gorkaio/gboy/pkg/cart"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

const testfile string = "../../roms/10-print.gb"
const testtitle string = "10 PRINT"

func TestLoadsFile(t *testing.T) {
	c, err := cart.New()
	assert.NoError(t, err)
	err = c.Load(testfile)
	assert.NoError(t, err)
	assert.Equal(t, c.Filename, testfile)
}

func TestReadsCartTitle(t *testing.T) {
	c, err := cart.New()
	assert.NoError(t, err)
	err = c.Load(testfile)
	assert.NoError(t, err)
	assert.Equal(t, c.Title, testtitle)
}

func TestReadsControllerType(t *testing.T) {
	c, err := cart.New()
	assert.NoError(t, err)
	err = c.Load(testfile)
	assert.NoError(t, err)
	assert.Equal(t, c.Type, cart.Type{ID: 0, Name: "MBC0", Description: "ROM only"})
}

func TestReadsMemory(t *testing.T) {
	c, err := cart.New()
	assert.NoError(t, err)
	err = c.Load(testfile)
	assert.NoError(t, err)
	assert.Equal(t, c.Read(0x137), byte('P'))
}
