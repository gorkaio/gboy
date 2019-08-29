package registers_test

import (
	registers "github.com/gorkaio/gboy/pkg/registers"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestByteRegistersUpdateTheirValues(t *testing.T) {
	reg := registers.NewByteRegister()
	reg.Set(0xCA)
	assert.Equal(t, byte(0xCA), reg.Get())
}

func TestMaskedByteRegistersDontSetMaskedBits(t *testing.T) {
	reg := registers.NewMaskedByteRegister(0xAA)
	reg.Set(0xF3)
	assert.Equal(t, byte(0xA2), reg.Get())
}

func TestMasksCanBeAppliedDynamicallyToByteRegisters(t *testing.T) {
	reg := registers.NewByteRegister()
	reg.Set(0xCA)
	assert.Equal(t, byte(0xCA), reg.Get())
	reg.SetMask(0xAA)
	assert.Equal(t, byte(0x8A), reg.Get())
	reg.SetMask(0xFF)
	assert.Equal(t, byte(0xCA), reg.Get())
	reg.SetMask(0xAA)
	reg.Set(0xCA)
	assert.Equal(t, byte(0x8A), reg.Get())
}

func TestByteRegistersIncreaseTheirValues(t *testing.T) {
	reg := registers.NewByteRegister()
	reg.Set(0xCA)
	reg.Inc()
	assert.Equal(t, byte(0xCB), reg.Get())
	reg.IncBy(2)
	assert.Equal(t, byte(0xCD), reg.Get())
}

func TestByteRegistersOverflowWithIncrease(t *testing.T) {
	reg := registers.NewByteRegister()
	reg.Set(0xFE)
	reg.IncBy(3)
	assert.Equal(t, byte(0x01), reg.Get())
}

func TestMaskedByteRegistersIncreaseRespectingTheirMask(t *testing.T) {
	reg := registers.NewMaskedByteRegister(0xAA)
	reg.Set(0xCA)
	assert.Equal(t, byte(0x8A), reg.Get())
	reg.Inc()
	assert.Equal(t, byte(0x8A), reg.Get())
	reg.IncBy(2)
	assert.Equal(t, byte(0x88), reg.Get())
}

func TestByteRegistersDecreaseTheirValues(t *testing.T) {
	reg := registers.NewByteRegister()
	reg.Set(0xCA)
	reg.Dec()
	assert.Equal(t, byte(0xC9), reg.Get())
	reg.DecBy(2)
	assert.Equal(t, byte(0xC7), reg.Get())
}

func TestMaskedByteRegistersDecreaseRespectingTheirMask(t *testing.T) {
	reg := registers.NewMaskedByteRegister(0xAA)
	reg.Set(0xCA)
	assert.Equal(t, byte(0x8A), reg.Get())
	reg.Dec()
	assert.Equal(t, byte(0x88), reg.Get())
	reg.DecBy(2)
	assert.Equal(t, byte(0x82), reg.Get())
}

func TestByteRegistersOverflowWithDecrease(t *testing.T) {
	reg := registers.NewByteRegister()
	reg.Set(0x01)
	reg.DecBy(3)
	assert.Equal(t, byte(0xFE), reg.Get())
}

func TestWordRegistersUpdateTheirValues(t *testing.T) {
	reg := registers.NewWordRegister()
	reg.Set(0xCAFE)
	assert.Equal(t, byte(0xCA), reg.H().Get())
	assert.Equal(t, byte(0xFE), reg.L().Get())
	assert.Equal(t, uint16(0xCAFE), reg.Get())
}

func TestMaskedWordRegistersDontSetMaskedBits(t *testing.T) {
	reg := registers.NewMaskedWordRegister(0xAAAA)
	reg.Set(0xF33F)
	assert.Equal(t, uint16(0xA22A), reg.Get())
}

func TestMasksCanBeAppliedDynamicallyToWordRegisters(t *testing.T) {
	reg := registers.NewWordRegister()
	reg.Set(0xCAFE)
	assert.Equal(t, uint16(0xCAFE), reg.Get())
	reg.SetMask(0xABBA)
	assert.Equal(t, uint16(0x8ABA), reg.Get())
	reg.SetMask(0xFFFF)
	assert.Equal(t, uint16(0xCAFE), reg.Get())
	reg.SetMask(0x9339)
	reg.Set(0x61CA)
	assert.Equal(t, uint16(0x0108), reg.Get())
}

func TestWordRegistersIncreaseTheirValues(t *testing.T) {
	reg := registers.NewWordRegister()
	reg.Set(0xCAFE)
	reg.Inc()
	assert.Equal(t, byte(0xCA), reg.H().Get())
	assert.Equal(t, byte(0xFF), reg.L().Get())
	assert.Equal(t, uint16(0xCAFF), reg.Get())
	reg.IncBy(2)
	assert.Equal(t, byte(0xCB), reg.H().Get())
	assert.Equal(t, byte(0x01), reg.L().Get())
	assert.Equal(t, uint16(0xCB01), reg.Get())
}

func TestMaskedWordRegistersIncreaseTheirValuesRespectingTheirMask(t *testing.T) {
	reg := registers.NewMaskedWordRegister(0xAAAA)
	reg.Set(0xCAFE)
	reg.Inc()
	assert.Equal(t, byte(0x8A), reg.H().Get())
	assert.Equal(t, byte(0xAA), reg.L().Get())
	assert.Equal(t, uint16(0x8AAA), reg.Get())
	reg.IncBy(2)
	assert.Equal(t, byte(0x8A), reg.H().Get())
	assert.Equal(t, byte(0xA8), reg.L().Get())
	assert.Equal(t, uint16(0x8AA8), reg.Get())
}

func TestWordRegistersIncreaseTheirValuesWithOverflow(t *testing.T) {
	reg := registers.NewWordRegister()
	reg.Set(0xFFFE)
	reg.IncBy(3)
	assert.Equal(t, byte(0x00), reg.H().Get())
	assert.Equal(t, byte(0x01), reg.L().Get())
	assert.Equal(t, uint16(0x0001), reg.Get())
}

func TestWordRegistersDecreaseTheirValues(t *testing.T) {
	reg := registers.NewWordRegister()
	reg.Set(0xCAFE)
	reg.Dec()
	assert.Equal(t, byte(0xCA), reg.H().Get())
	assert.Equal(t, byte(0xFD), reg.L().Get())
	assert.Equal(t, uint16(0xCAFD), reg.Get())
	reg.DecBy(0x102)
	assert.Equal(t, byte(0xC9), reg.H().Get())
	assert.Equal(t, byte(0xFB), reg.L().Get())
	assert.Equal(t, uint16(0xC9FB), reg.Get())
}

func TestWordRegistersDecreaseTheirValuesWithOverflow(t *testing.T) {
	reg := registers.NewWordRegister()
	reg.Set(0x0001)
	reg.DecBy(3)
	assert.Equal(t, byte(0xFF), reg.H().Get())
	assert.Equal(t, byte(0xFE), reg.L().Get())
	assert.Equal(t, uint16(0xFFFE), reg.Get())
}
