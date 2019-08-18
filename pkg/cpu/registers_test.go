package cpu_test

import (
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestByteRegistersUpdateTheirValues(t *testing.T) {
	r := cpu.ByteRegister{}
	r.Set(0xFE)
	assert.Equal(t, uint8(0xFE), r.Get())
}

func TestByteRegistersCanIncreaseTheirValue(t *testing.T) {
	r := cpu.ByteRegister{}
	r.Set(0xFE)
	r.Inc()
	assert.Equal(t, uint8(0xFF), r.Get())
}

func TestByteRegistersCanIncreaseTheirValueWithOverflow(t *testing.T) {
	r := cpu.ByteRegister{}
	r.Set(0xFE)
	r.IncBy(3)
	assert.Equal(t, uint8(0x01), r.Get())
}

func TestByteRegistersCanDecreaseTheirValue(t *testing.T) {
	r := cpu.ByteRegister{}
	r.Set(0xFE)
	r.Dec()
	assert.Equal(t, uint8(0xFD), r.Get())
}

func TestByteRegistersCanDecreaseTheirValueWithOverflow(t *testing.T) {
	r := cpu.ByteRegister{}
	r.Set(0x02)
	r.DecBy(3)
	assert.Equal(t, uint8(0xFF), r.Get())
}


func TestWordRegistersUpdateTheirValues(t *testing.T) {
	r := cpu.WordRegister{}
	r.Set(0xFEFE)
	assert.Equal(t, uint16(0xFEFE), r.Get())
}

func TestWordRegistersCanIncreaseTheirValue(t *testing.T) {
	r := cpu.WordRegister{}
	r.Set(0xFFFE)
	r.Inc()
	assert.Equal(t, uint16(0xFFFF), r.Get())
}

func TestWordRegistersCanIncreaseTheirValueWithOverflow(t *testing.T) {
	r := cpu.WordRegister{}
	r.Set(0xFFFE)
	r.IncBy(3)
	assert.Equal(t, uint16(0x0001), r.Get())
}

func TestWordRegistersCanDecreaseTheirValue(t *testing.T) {
	r := cpu.WordRegister{}
	r.Set(0xFFFE)
	r.Dec()
	assert.Equal(t, uint16(0xFFFD), r.Get())
}

func TestWordRegistersCanDecreaseTheirValueWithOverflow(t *testing.T) {
	r := cpu.WordRegister{}
	r.Set(0x0002)
	r.DecBy(3)
	assert.Equal(t, uint16(0xFFFF), r.Get())
}
