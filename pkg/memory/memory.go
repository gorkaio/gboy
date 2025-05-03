package memory

import (
	"github.com/gorkaio/gboy/pkg/video"
)

//go:generate mockgen -destination=mocks/cart_mock.go -package=memory_mock github.com/gorkaio/gboy/pkg/memory Cart

const cartAddressHigh = 0x7FFF

// Cart interface for the cart
type Cart interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}

// Memory defines the memory structure
type Memory struct {
	cart       Cart
	system     []byte
	cartLoaded bool
	gpu        *video.GPU
}

// New creates a new memory
func New(display video.Display) *Memory {
	mem := Memory{
		system:     make([]byte, 0x8000),
		cartLoaded: false,
	}
	
	// Initialize the GPU with the display
	mem.gpu = video.NewGPU(display)
	
	return &mem
}

// Eject ejects the current cartdrige
func (mem *Memory) Eject() {
	mem.cart = nil
	mem.cartLoaded = false
}

// Load loads a cart from file
func (mem *Memory) Load(cart Cart) {
	mem.cart = cart
	mem.cartLoaded = true
}

func (mem *Memory) Read(address uint16) byte {
	if addressInCart(address) {
		if mem.cartLoaded {
			return mem.cart.Read(address)
		}
		return 0xFF
	}
	
	// Check if the address is in the GPU range (0x8000-0x9FFF for VRAM, 0xFE00-0xFE9F for OAM, 0xFF40-0xFF4B for GPU registers)
	if (address >= 0x8000 && address <= 0x9FFF) || (address >= 0xFE00 && address <= 0xFE9F) || (address >= 0xFF40 && address <= 0xFF4B) {
		return mem.gpu.Read(address)
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
	
	// Check if the address is in the GPU range (0x8000-0x9FFF for VRAM, 0xFE00-0xFE9F for OAM, 0xFF40-0xFF4B for GPU registers)
	if (address >= 0x8000 && address <= 0x9FFF) || (address >= 0xFE00 && address <= 0xFE9F) || (address >= 0xFF40 && address <= 0xFF4B) {
		mem.gpu.Write(address, data)
		return
	}

	mem.system[address&0x7FFF] = data
}

func addressInCart(address uint16) bool {
	return (address <= cartAddressHigh)
}

// GetGPU returns the GPU instance
func (mem *Memory) GetGPU() *video.GPU {
	return mem.gpu
}
