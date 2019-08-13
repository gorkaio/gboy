package memory

import (
	cart "github.com/gorkaio/gboy/pkg/cart"
)

const cartAddressHigh = 0x7FFF

// Memory defines the memory structure
type Memory struct {
	cart cart.CartController
}

// NewMemory creates a new memory
func NewMemory(cart cart.CartController) (*Memory, error) {
	mem := Memory{
		cart: cart,
	}
	return &mem, nil
}

func (mem *Memory) Read(address uint16) byte {
	if (address <= cartAddressHigh) {
		return mem.cart.Read(address)
	}
	return 0
}
