package cart_test

import (
	cart "github.com/gorkaio/gboy/pkg/cart"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

const testfile string = "../../roms/10-print.gb"

func TestFailsWhenFileCannotBeLoaded(t *testing.T) {
	loader := cart.NewFileLoader()
	_, err := loader.Load("thisdoesnotexist.gb")
	assert.Error(t, err)
}

func TestFailsWhenFileIsNotGBRom(t *testing.T) {
	loader := cart.NewFileLoader()
	_, err := loader.Load("../../README.md")
	assert.Error(t, err)
}

func TestSucceedsWhenFileIsGBRomAndCanBeLoaded(t *testing.T) {
	loader := cart.NewFileLoader()
	_, err := loader.Load(testfile)
	assert.NoError(t, err)
}
