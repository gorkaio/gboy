package memory

import (
	assert "gotest.tools/assert"
	testing "testing"
	cart "github.com/gorkaio/gboy/pkg/cart"
	gomock "github.com/golang/mock/gomock"
	rand "math/rand"
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
	if (err != nil) {
		t.Error("Unable to initialise memory")
	}

	assert.Equal(t, mem.Read(address), data)
}