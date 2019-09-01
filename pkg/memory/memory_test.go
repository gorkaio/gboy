package memory_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mocks "github.com/gorkaio/gboy/pkg/memory/mocks"
	"github.com/gorkaio/gboy/pkg/memory"
	"testing"
)

const (
	mockAddress = uint16(0x7234)
	mockData = byte(0xCA)
)

func TestReadsAddressInCartRange(t *testing.T) {
	address := uint16(0x7234)
	data := byte(0xCA)

	ctrlCart := gomock.NewController(t)
	defer ctrlCart.Finish()
	cart := mocks.NewMockCart(ctrlCart)
	cart.
		EXPECT().
		Read(address).
		Return(data)

	mem := memory.New()
	mem.Load(cart)

	assert.Equal(t, data, mem.Read(address))
}

func TestWritesAddressInCartRange(t *testing.T) {
	address := uint16(0x7FFF)
	data := byte(0xFF)

	ctrlCart := gomock.NewController(t)
	defer ctrlCart.Finish()
	cart := mocks.NewMockCart(ctrlCart)
	cart.
		EXPECT().
		Write(address, data)

	mem := memory.New()
	mem.Load(cart)

	mem.Write(address, data)
}

func TestWorksWithAddressInSystemRange(t *testing.T) {
	address := uint16(0x8000)
	data := byte(0xFE)

	mem := memory.New()
	mem.Write(address, data)
	assert.Equal(t, data, mem.Read(address))
}

func TestEjectsCart(t *testing.T) {
	address := uint16(0x7234)
	data := byte(0xCA)

	ctrlCart := gomock.NewController(t)
	defer ctrlCart.Finish()
	cart := mocks.NewMockCart(ctrlCart)
	cart.
		EXPECT().
		Read(address).
		Return(data)

	mem := memory.New()
	mem.Load(cart)

	assert.Equal(t, data, mem.Read(address))

	mem.Eject()
	assert.Equal(t, byte(0xFF), mem.Read(address))
}
