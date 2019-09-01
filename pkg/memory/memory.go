package memory

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
}

// New creates a new memory
func New() *Memory {
	mem := Memory{
		system:     make([]byte, 0x8000),
		cartLoaded: false,
	}
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
