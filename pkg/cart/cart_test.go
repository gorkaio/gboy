package cart_test

import (
	"github.com/gorkaio/gboy/pkg/cart"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadsCartTitle(t *testing.T) {
	c, err := loadTestCart()
	assert.NoError(t, err)
	assert.Equal(t, c.Title(), "10 PRINT")
}

func TestReadsControllerType(t *testing.T) {
	c, err := loadTestCart()
	assert.NoError(t, err)
	assert.Equal(t, c.Type(), cart.Type{ID: 0, Name: "MBC0", Description: "ROM only"})
}

func TestReadsMemory(t *testing.T) {
	c, err := loadTestCart()
	assert.NoError(t, err)
	assert.Equal(t, c.Read(0x137), byte('P'))
}

func loadTestCart() (cart.CartInterface, error) {
	loader := cart.NewFileLoader()
	c, err := loader.Load(testfile)
	return c, err
}
