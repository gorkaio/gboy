package cart

//go:generate mockgen -destination=../mocks/mock_cart.go -package=mocks github.com/gorkaio/gboy/pkg/cart CartInterface
//go:generate mockgen -destination=../mocks/mock_loader.go -package=mocks github.com/gorkaio/gboy/pkg/cart LoaderInterface
//go:generate mockgen -destination=../mocks/mock_mbc.go -package=mocks github.com/gorkaio/gboy/pkg/cart MemoryBankControllerInterface

import (
	"errors"
	"fmt"
	"strings"
)

const titleStartAddr, titleEndAddr = 0x134, 0x143
const cartTypeAddr = 0x147

// CartInterface interface for the cart
type CartInterface interface {
	Title() string
	Type() Type
	MemoryBankControllerInterface
}

// MemoryBankControllerInterface interface for the MBC
type MemoryBankControllerInterface interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}

// Type defines the cartdrige type
type Type struct {
	ID          byte
	Name        string
	Description string
}

// Cart contains the cartdridge data
type Cart struct {
	controller MemoryBankControllerInterface
	CartInterface
}

func (cart *Cart) Read(address uint16) byte {
	return cart.controller.Read(address)
}

func (cart *Cart) Write(address uint16, data uint8) {
	cart.controller.Write(address, data)
}

func (cart *Cart) Title() string {
	title := ""
	for address := uint16(titleStartAddr); address < titleEndAddr; address++ {
		chr := cart.controller.Read(address)
		if chr != 0x00 {
			title += string(chr)
		}
	}
	return strings.TrimSpace(title)
}

func (cart *Cart) Type() Type {
	cartTypeID := cart.controller.Read(cartTypeAddr)
	if cartTypeID == 0 {
		return Type{ID: cartTypeID, Name: "MBC0", Description: "ROM only"}
	}
	return Type{ID: cartTypeID, Name: "UNKNOWN", Description: "Unknown"}
}

// New loads cartdridge information from file
func newCart(data []byte) (*Cart, error) {
	cartTypeID := data[cartTypeAddr]
	if cartTypeID != 0 {
		msg := fmt.Sprintf("Unknown memory controller (%#02x). Cannot load ROM.", cartTypeID)
		return nil, errors.New(msg)
	}

	cart := &Cart{
		controller: NewMBC0(data),
	}
	return cart, nil
}
