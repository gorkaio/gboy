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
