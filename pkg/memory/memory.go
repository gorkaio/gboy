package memory

//go:generate mockgen -destination=../mocks/mock_memory.go -package=mocks github.com/gorkaio/gboy/pkg/memory MemoryInterface

import (
	cart "github.com/gorkaio/gboy/pkg/cart"
)

const cartAddressHigh = 0x7FFF

type MemoryInterface interface {
	Read(address uint16) byte
	Write(address uint16, data byte)
}

// Memory defines the memory structure
type Memory struct {
	cart cart.CartInterface
}

// New creates a new memory
func New(cart cart.CartInterface) (*Memory, error) {
	mem := Memory{
		cart: cart,
	}
	return &mem, nil
}

// LoadRomFile loads a cart from file
func (mem *Memory) LoadRomFile(romfile string) error {
	err := mem.cart.Load(romfile)
	if (err != nil) {
		return err
	}

	return nil
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
