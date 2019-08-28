package memory

//go:generate mockgen -destination=../mocks/mock_memory.go -package=mocks github.com/gorkaio/gboy/pkg/memory MemoryInterface

import (
	"github.com/gorkaio/gboy/pkg/cart"
)

const cartAddressHigh = 0x7FFF

type MemoryInterface interface {
	Load(romfile string) error
	Eject()
	Read(address uint16) byte
	Write(address uint16, data byte)
}

// Memory defines the memory structure
type Memory struct {
	cart       cart.CartInterface
	loader     cart.Loader
	system     []byte
	cartLoaded bool
}

// New creates a new memory
func New(loader cart.Loader) (MemoryInterface, error) {
	mem := Memory{
		system:     make([]byte, 0x8000),
		loader:     loader,
		cartLoaded: false,
	}
	return &mem, nil
}

// Eject ejects the current cartdrige
func (mem *Memory) Eject() {
	mem.cart = nil
	mem.cartLoaded = false
}

// Load loads a cart from file
func (mem *Memory) Load(romfile string) error {
	c, err := mem.loader.Load(romfile)
	if err != nil {
		return err
	}
	mem.cart = c
	mem.cartLoaded = true
	return nil
}

func (mem *Memory) Read(address uint16) byte {
	if addressInCart(address) {
		if mem.cartLoaded {
			return mem.cart.Read(address)
		}
		return 0xFF
	}

	return mem.system[address&0x7FFF]
}

func (mem *Memory) Write(address uint16, data byte) {
	if addressInCart(address) {
		if mem.cartLoaded {
			mem.cart.Write(address, data)
		}
		return
	}

	mem.system[address&0x7FFF] = data
}

func addressInCart(address uint16) bool {
	return (address <= cartAddressHigh)
}
