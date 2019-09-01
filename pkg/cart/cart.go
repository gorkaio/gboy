package cart

//go:generate mockgen -destination=mocks/mbc_mock.go -package=cart_mock github.com/gorkaio/gboy/pkg/cart MemoryBankController

import (
	"errors"
	"fmt"
	"strings"
)

const titleStartAddr, titleEndAddr = 0x134, 0x143
const cartTypeAddr = 0x147

// MemoryBankController interface for the MBC
type MemoryBankController interface {
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
	controller MemoryBankController
}

func (cart *Cart) Read(address uint16) byte {
	return cart.controller.Read(address)
}

func (cart *Cart) Write(address uint16, data byte) {
	cart.controller.Write(address, data)
}

// Title gets title for the cartdrige
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

// Type gets cartdrige type
func (cart *Cart) Type() Type {
	cartTypeID := cart.controller.Read(cartTypeAddr)
	if cartTypeID == 0 {
		return Type{ID: cartTypeID, Name: "MBC0", Description: "ROM only"}
	}
	return Type{ID: cartTypeID, Name: "UNKNOWN", Description: "Unknown"}
}

// NewCart loads cartdridge information from file
func NewCart(data []byte) (*Cart, error) {
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
