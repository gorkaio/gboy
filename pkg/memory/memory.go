package memory

import (
	cart "github.com/gorkaio/gboy/pkg/cart"
)

const cartAddressHigh = 0x7FFF

type Controller interface {
	Read(address uint16) byte
	Write(address uint16, data byte)
}

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
	if addressInCart(address) {
		return mem.cart.Read(address)
	}
	return 0
}

func (mem *Memory) Write(address uint16, data byte) {
	if addressInCart(address) {
		mem.cart.Write(address, data)
	}
}

func addressInCart(address uint16) bool {
	return (address <= cartAddressHigh)
}
