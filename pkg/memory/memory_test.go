package memory

import (
	gomock "github.com/golang/mock/gomock"
	cart "github.com/gorkaio/gboy/pkg/cart"
	assert "gotest.tools/assert"
	rand "math/rand"
	testing "testing"
)

func TestReadsAddressInCartRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	address := uint16(rand.Intn(0x7FFF))
	data := byte(rand.Intn(0xFF))

	cart := cart.NewMockCartController(ctrl)
	cart.
		EXPECT().
		Read(address).
		Return(data).
		AnyTimes()

	mem, err := NewMemory(cart)
	if err != nil {
		t.Error("Unable to initialise memory")
	}

	assert.Equal(t, mem.Read(address), data)
}

func TestWritesAddressInCartRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	address := uint16(rand.Intn(0x7FFF))
	data := byte(rand.Intn(0xFF))

	cart := cart.NewMockCartController(ctrl)
	cart.
		EXPECT().
		Write(address, data)

	mem, err := NewMemory(cart)
	if err != nil {
		t.Error("Unable to initialise memory")
	}

	mem.Write(address, data)
}
