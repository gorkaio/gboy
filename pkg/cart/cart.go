package cart

//go:generate mockgen -destination=../mocks/mock_cart.go -package=mocks github.com/gorkaio/gboy/pkg/cart CartInterface
//go:generate mockgen -destination=../mocks/mock_mbc.go -package=mocks github.com/gorkaio/gboy/pkg/cart MemoryBankControllerInterface

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const titleStartAddr, titleEndAddr = 0x134, 0x143
const cartTypeAddr = 0x147

// CartInterface interface for the cart
type CartInterface interface {
	Load(romfile string) error
	MemoryBankControllerInterface
}

// MemoryBankControllerInterface interface for the MBC
type MemoryBankControllerInterface interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}

// Type defines the cartdrige type
type Type struct {
	ID          int
	Name        string
	Description string
}

// Cart contains the cartdridge data
type Cart struct {
	Filename   string
	Title      string
	Type       Type
	controller MemoryBankControllerInterface
	CartInterface
	MemoryBankControllerInterface
}

// New returns a new empty cart
func New() (*Cart, error) {
	return &Cart{}, nil
}

func (cart *Cart) Read(address uint16) byte {
	return cart.controller.Read(address)
}

func (cart *Cart) Write(address uint16, data uint8) {
	cart.controller.Write(address, data)
}

func title(data *[]byte) string {
	title := string((*data)[titleStartAddr:titleEndAddr])
	return strings.Trim(title, "\x00")
}

func cartType(data *[]byte) Type {
	cartTypeID := int((*data)[cartTypeAddr])
	if cartTypeID == 0 {
		return Type{ID: cartTypeID, Name: "MBC0", Description: "ROM only"}
	}
	return Type{ID: cartTypeID, Name: "UNKNOWN", Description: "Unknown"}
}

// Load loads cartdridge information from file
func (cart *Cart) Load(romfile string) error {
	data, err := ioutil.ReadFile(romfile)
	if err != nil {
		return err
	}

	cartType := cartType(&data)
	if cartType.ID != 0 {
		msg := fmt.Sprintf("Unknown memory controller (%#02x). Cannot load ROM.", cartType.ID)
		return errors.New(msg)
	}

	cart.Filename = romfile
	cart.Title = title(&data)
	cart.Type = cartType
	cart.controller = NewMBC0(data)

	return nil
}
