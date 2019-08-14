package cart

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const titleStartAddr, titleEndAddr = 0x134, 0x143
const cartTypeAddr = 0x147

// Controller defines the interface for accessing the cart
type Controller interface {
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
	controller Controller
}

func new(data []byte, filename string) (*Cart, error) {
	cartType := cartType(&data)
	if cartType.ID != 0 {
		msg := fmt.Sprintf("Unknown memory controller (%#02x). Cannot load ROM.", cartType.ID)
		return nil, errors.New(msg)
	}

	cart := Cart{
		Filename:   filename,
		Title:      title(&data),
		Type:       cartType,
		controller: NewMBC0(data),
	}

	return &cart, nil
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

// LoadFromFile loads cartdridge information from file
func LoadFromFile(filename string) (*Cart, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cart, err := new(data, filename)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
