package cart

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const titleStartAddr, titleEndAddr = 0x134, 0x143
const cartTypeAddr = 0x147

type CartController interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}

type CartType struct {
	id   int
	name string
}

type Cart struct {
	filename string
	title    string
	cartType CartType
	CartController
}

func newCart(data []byte, filename string) (*Cart, error) {
	cartType := cartType(&data)
	if cartType.id != 0 {
		msg := fmt.Sprintf("Unknown memory controller (%#02x). Cannot load ROM.", cartType.id)
		return nil, errors.New(msg)
	}
	cart := Cart{
		filename:       filename,
		title:          title(&data),
		cartType:       cartType,
		CartController: NewMBC0(data),
	}
	return &cart, nil
}

func (c *Cart) Title() string {
	return c.title
}

func (c *Cart) Type() string {
	return fmt.Sprintf("%s (%#02x)", c.cartType.name, c.cartType.id)
}

func title(data *[]byte) string {
	title := string((*data)[titleStartAddr:titleEndAddr])
	return strings.Trim(title, "\x00")
}

func cartType(data *[]byte) CartType {
	cartTypeId := int((*data)[cartTypeAddr])
	if cartTypeId == 0 {
		return CartType{id: cartTypeId, name: "ROM only"}
	}
	return CartType{id: cartTypeId, name: "UNKNOWN"}
}

func LoadFromFile(filename string) (*Cart, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cart, err := newCart(data, filename)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
