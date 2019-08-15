package memory_test

import (
	gomock "github.com/golang/mock/gomock"
	mocks "github.com/gorkaio/gboy/pkg/mocks"
	memory "github.com/gorkaio/gboy/pkg/memory"
	assert "github.com/stretchr/testify/assert"
	rand "math/rand"
	testing "testing"
)

func TestReadsAddressInCartRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	address := uint16(rand.Intn(0x7FFF))
	data := byte(rand.Intn(0xFF))

	cart := mocks.NewMockCartInterface(ctrl)
	cart.
		EXPECT().
		Read(address).
		Return(data).
		AnyTimes()

	mem, err := memory.New(cart)
	assert.NoError(t, err)
	assert.Equal(t, mem.Read(address), data)
}

func TestWritesAddressInCartRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	address := uint16(rand.Intn(0x7FFF))
	data := byte(rand.Intn(0xFF))

	cart := mocks.NewMockCartInterface(ctrl)
	cart.
		EXPECT().
		Write(address, data)

	mem, err := memory.New(cart)
	assert.NoError(t, err)
	mem.Write(address, data)
}

func TestLoadsCartFromFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	romfile := "rom.gb"
	cartdrige := mocks.NewMockCartInterface(ctrl)
	cartdrige.
		EXPECT().
		Load(romfile)
	
	mem, err := memory.New(cartdrige)
	assert.NoError(t, err)

	err = mem.LoadRomFile(romfile)
	assert.NoError(t, err)
}
