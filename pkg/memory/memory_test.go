package memory_test

import (
	"github.com/golang/mock/gomock"
	"github.com/gorkaio/gboy/pkg/memory"
	"github.com/gorkaio/gboy/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadsAddressInCartRange(t *testing.T) {
	romfile := "romfile.gb"
	address := uint16(0x7234)
	data := byte(0xCA)

	ctrlCart := gomock.NewController(t)
	defer ctrlCart.Finish()
	cart := mocks.NewMockCartInterface(ctrlCart)
	cart.
		EXPECT().
		Read(address).
		Return(data)

	ctrlLoader := gomock.NewController(t)
	defer ctrlLoader.Finish()
	loader := mocks.NewMockLoaderInterface(ctrlLoader)
	loader.
		EXPECT().
		Load(romfile).
		Return(cart, nil)

	mem, err := memory.New(loader)
	assert.NoError(t, err)

	err = mem.Load(romfile)
	assert.NoError(t, err)

	assert.Equal(t, data, mem.Read(address))
}

func TestWritesAddressInCartRange(t *testing.T) {
	address := uint16(0x7FFF)
	data := byte(0xFF)
	romfile := "romfile.gb"

	ctrlCart := gomock.NewController(t)
	defer ctrlCart.Finish()
	cart := mocks.NewMockCartInterface(ctrlCart)
	cart.
		EXPECT().
		Write(address, data)

	ctrlLoader := gomock.NewController(t)
	defer ctrlLoader.Finish()
	loader := mocks.NewMockLoaderInterface(ctrlLoader)
	loader.
		EXPECT().
		Load(romfile).
		Return(cart, nil)

	mem, err := memory.New(loader)
	assert.NoError(t, err)

	err = mem.Load(romfile)
	assert.NoError(t, err)

	mem.Write(address, data)
}

func TestWorksWithAddressInSystemRange(t *testing.T) {
	address := uint16(0x8000)
	data := byte(0xFE)

	ctrlLoader := gomock.NewController(t)
	defer ctrlLoader.Finish()
	loader := mocks.NewMockLoaderInterface(ctrlLoader)

	mem, err := memory.New(loader)
	assert.NoError(t, err)

	mem.Write(address, data)
	assert.Equal(t, data, mem.Read(address))
}

func TestEjectsCart(t *testing.T) {
	romfile := "romfile.gb"
	address := uint16(0x7234)
	data := byte(0xCA)

	ctrlCart := gomock.NewController(t)
	defer ctrlCart.Finish()
	cart := mocks.NewMockCartInterface(ctrlCart)
	cart.
		EXPECT().
		Read(address).
		Return(data)

	ctrlLoader := gomock.NewController(t)
	defer ctrlLoader.Finish()
	loader := mocks.NewMockLoaderInterface(ctrlLoader)
	loader.
		EXPECT().
		Load(romfile).
		Return(cart, nil)

	mem, err := memory.New(loader)
	assert.NoError(t, err)

	err = mem.Load(romfile)
	assert.NoError(t, err)

	assert.Equal(t, data, mem.Read(address))

	mem.Eject()
	assert.Equal(t, byte(0xFF), mem.Read(address))
}
