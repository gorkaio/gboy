package cpu_test

import (
	gomock "github.com/golang/mock/gomock"
	mocks "github.com/gorkaio/gboy/pkg/mocks"
	cpu "github.com/gorkaio/gboy/pkg/cpu"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestByteRegistersUpdateTheirValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.A.Set(0xCA)
	c.F.Set(0xFE)
	assert.Equal(t, uint8(0xCA), c.A.Get())
	assert.Equal(t, uint8(0xFE), c.F.Get())
	assert.Equal(t, uint16(0xCAFE), c.AF.Get())

	c.B.Set(0xCA)
	c.C.Set(0xFE)
	assert.Equal(t, uint8(0xCA), c.B.Get())
	assert.Equal(t, uint8(0xFE), c.C.Get())
	assert.Equal(t, uint16(0xCAFE), c.BC.Get())

	c.D.Set(0xCA)
	c.E.Set(0xFE)
	assert.Equal(t, uint8(0xCA), c.D.Get())
	assert.Equal(t, uint8(0xFE), c.E.Get())
	assert.Equal(t, uint16(0xCAFE), c.DE.Get())

	c.H.Set(0xCA)
	c.L.Set(0xFE)
	assert.Equal(t, uint8(0xCA), c.H.Get())
	assert.Equal(t, uint8(0xFE), c.L.Get())
	assert.Equal(t, uint16(0xCAFE), c.HL.Get())
}

func TestByteRegistersCanIncreaseTheirValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.A.Set(0xCA)
	c.F.Set(0xFE)
	c.A.Inc()
	c.F.Inc()
	assert.Equal(t, uint8(0xCB), c.A.Get())
	assert.Equal(t, uint8(0xFF), c.F.Get())
	assert.Equal(t, uint16(0xCBFF), c.AF.Get())

	c.B.Set(0xCA)
	c.C.Set(0xFE)
	c.B.Inc()
	c.C.Inc()
	assert.Equal(t, uint8(0xCB), c.B.Get())
	assert.Equal(t, uint8(0xFF), c.C.Get())
	assert.Equal(t, uint16(0xCBFF), c.BC.Get())

	c.D.Set(0xCA)
	c.E.Set(0xFE)
	c.D.Inc()
	c.E.Inc()
	assert.Equal(t, uint8(0xCB), c.D.Get())
	assert.Equal(t, uint8(0xFF), c.E.Get())
	assert.Equal(t, uint16(0xCBFF), c.DE.Get())

	c.H.Set(0xCA)
	c.L.Set(0xFE)
	c.H.Inc()
	c.L.Inc()
	assert.Equal(t, uint8(0xCB), c.H.Get())
	assert.Equal(t, uint8(0xFF), c.L.Get())
	assert.Equal(t, uint16(0xCBFF), c.HL.Get())
}

func TestByteRegistersCanIncreaseTheirValueWithOverflow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.A.Set(0xFE)
	c.F.Set(0xFF)
	c.A.IncBy(3)
	c.F.IncBy(3)
	assert.Equal(t, uint8(0x01), c.A.Get())
	assert.Equal(t, uint8(0x02), c.F.Get())
	assert.Equal(t, uint16(0x0102), c.AF.Get())

	c.B.Set(0xFE)
	c.C.Set(0xFF)
	c.B.IncBy(3)
	c.C.IncBy(3)
	assert.Equal(t, uint8(0x01), c.B.Get())
	assert.Equal(t, uint8(0x02), c.C.Get())
	assert.Equal(t, uint16(0x0102), c.BC.Get())

	c.D.Set(0xFE)
	c.E.Set(0xFF)
	c.D.IncBy(3)
	c.E.IncBy(3)
	assert.Equal(t, uint8(0x01), c.D.Get())
	assert.Equal(t, uint8(0x02), c.E.Get())
	assert.Equal(t, uint16(0x0102), c.DE.Get())

	c.H.Set(0xFE)
	c.L.Set(0xFF)
	c.H.IncBy(3)
	c.L.IncBy(3)
	assert.Equal(t, uint8(0x01), c.H.Get())
	assert.Equal(t, uint8(0x02), c.L.Get())
	assert.Equal(t, uint16(0x0102), c.HL.Get())
}

func TestByteRegistersCanDecreaseTheirValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.A.Set(0xCA)
	c.F.Set(0xFE)
	c.A.Dec()
	c.F.Dec()
	assert.Equal(t, uint8(0xC9), c.A.Get())
	assert.Equal(t, uint8(0xFD), c.F.Get())
	assert.Equal(t, uint16(0xC9FD), c.AF.Get())

	c.B.Set(0xCA)
	c.C.Set(0xFE)
	c.B.Dec()
	c.C.Dec()
	assert.Equal(t, uint8(0xC9), c.B.Get())
	assert.Equal(t, uint8(0xFD), c.C.Get())
	assert.Equal(t, uint16(0xC9FD), c.BC.Get())

	c.D.Set(0xCA)
	c.E.Set(0xFE)
	c.D.Dec()
	c.E.Dec()
	assert.Equal(t, uint8(0xC9), c.D.Get())
	assert.Equal(t, uint8(0xFD), c.E.Get())
	assert.Equal(t, uint16(0xC9FD), c.DE.Get())

	c.H.Set(0xCA)
	c.L.Set(0xFE)
	c.H.Dec()
	c.L.Dec()
	assert.Equal(t, uint8(0xC9), c.H.Get())
	assert.Equal(t, uint8(0xFD), c.L.Get())
	assert.Equal(t, uint16(0xC9FD), c.HL.Get())
}

func TestByteRegistersCanDecreaseTheirValueWithOverflow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.A.Set(0x01)
	c.F.Set(0x02)
	c.A.DecBy(3)
	c.F.DecBy(3)
	assert.Equal(t, uint8(0xFE), c.A.Get())
	assert.Equal(t, uint8(0xFF), c.F.Get())
	assert.Equal(t, uint16(0xFEFF), c.AF.Get())

	c.B.Set(0x01)
	c.C.Set(0x02)
	c.B.DecBy(3)
	c.C.DecBy(3)
	assert.Equal(t, uint8(0xFE), c.B.Get())
	assert.Equal(t, uint8(0xFF), c.C.Get())
	assert.Equal(t, uint16(0xFEFF), c.BC.Get())

	c.D.Set(0x01)
	c.E.Set(0x02)
	c.D.DecBy(3)
	c.E.DecBy(3)
	assert.Equal(t, uint8(0xFE), c.D.Get())
	assert.Equal(t, uint8(0xFF), c.E.Get())
	assert.Equal(t, uint16(0xFEFF), c.DE.Get())

	c.H.Set(0x01)
	c.L.Set(0x02)
	c.H.DecBy(3)
	c.L.DecBy(3)
	assert.Equal(t, uint8(0xFE), c.H.Get())
	assert.Equal(t, uint8(0xFF), c.L.Get())
	assert.Equal(t, uint16(0xFEFF), c.HL.Get())
}

func TestWordRegistersUpdateTheirValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.AF.Set(0xCAFE)
	assert.Equal(t, uint16(0xCAFE), c.AF.Get())
	assert.Equal(t, uint8(0xCA), c.A.Get())
	assert.Equal(t, uint8(0xFE), c.F.Get())
	
	c.BC.Set(0xCAFE)
	assert.Equal(t, uint16(0xCAFE), c.BC.Get())
	assert.Equal(t, uint8(0xCA), c.B.Get())
	assert.Equal(t, uint8(0xFE), c.C.Get())

	c.DE.Set(0xCAFE)
	assert.Equal(t, uint16(0xCAFE), c.DE.Get())
	assert.Equal(t, uint8(0xCA), c.D.Get())
	assert.Equal(t, uint8(0xFE), c.E.Get())

	c.HL.Set(0xCAFE)
	assert.Equal(t, uint16(0xCAFE), c.HL.Get())
	assert.Equal(t, uint8(0xCA), c.H.Get())
	assert.Equal(t, uint8(0xFE), c.L.Get())

	c.SP.Set(0xFEFE)
	assert.Equal(t, uint16(0xFEFE), c.SP.Get())

	c.PC.Set(0xFEFE)
	assert.Equal(t, uint16(0xFEFE), c.PC.Get())
}

func TestWordRegistersCanIncreaseTheirValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.AF.Set(0xCAFE)
	c.AF.Inc()
	assert.Equal(t, uint16(0xCAFF), c.AF.Get())
	assert.Equal(t, uint8(0xCA), c.A.Get())
	assert.Equal(t, uint8(0xFF), c.F.Get())

	c.BC.Set(0xCAFE)
	c.BC.Inc()
	assert.Equal(t, uint16(0xCAFF), c.BC.Get())
	assert.Equal(t, uint8(0xCA), c.B.Get())
	assert.Equal(t, uint8(0xFF), c.C.Get())

	c.DE.Set(0xCAFE)
	c.DE.Inc()
	assert.Equal(t, uint16(0xCAFF), c.DE.Get())
	assert.Equal(t, uint8(0xCA), c.D.Get())
	assert.Equal(t, uint8(0xFF), c.E.Get())

	c.HL.Set(0xCAFE)
	c.HL.Inc()
	assert.Equal(t, uint16(0xCAFF), c.HL.Get())
	assert.Equal(t, uint8(0xCA), c.H.Get())
	assert.Equal(t, uint8(0xFF), c.L.Get())

	c.SP.Set(0xFEFE)
	c.SP.Inc()
	assert.Equal(t, uint16(0xFEFF), c.SP.Get())

	c.PC.Set(0xFEFE)
	c.PC.Inc()
	assert.Equal(t, uint16(0xFEFF), c.PC.Get())
}

func TestWordRegistersCanIncreaseTheirValueWithOverflow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.AF.Set(0xFFFE)
	c.AF.IncBy(3)
	assert.Equal(t, uint16(0x0001), c.AF.Get())
	assert.Equal(t, uint8(0x00), c.A.Get())
	assert.Equal(t, uint8(0x01), c.F.Get())

	c.BC.Set(0xFFFE)
	c.BC.IncBy(3)
	assert.Equal(t, uint16(0x0001), c.BC.Get())
	assert.Equal(t, uint8(0x00), c.B.Get())
	assert.Equal(t, uint8(0x01), c.C.Get())

	c.DE.Set(0xFFFE)
	c.DE.IncBy(3)
	assert.Equal(t, uint16(0x0001), c.DE.Get())
	assert.Equal(t, uint8(0x00), c.D.Get())
	assert.Equal(t, uint8(0x01), c.E.Get())

	c.HL.Set(0xFFFE)
	c.HL.IncBy(3)
	assert.Equal(t, uint16(0x0001), c.HL.Get())
	assert.Equal(t, uint8(0x00), c.H.Get())
	assert.Equal(t, uint8(0x01), c.L.Get())

	c.SP.Set(0xFFFE)
	c.SP.IncBy(3)
	assert.Equal(t, uint16(0x0001), c.SP.Get())

	c.PC.Set(0xFFFE)
	c.PC.IncBy(3)
	assert.Equal(t, uint16(0x0001), c.PC.Get())
}

func TestWordRegistersCanDecreaseTheirValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.AF.Set(0xCAFE)
	c.AF.Dec()
	assert.Equal(t, uint16(0xCAFD), c.AF.Get())
	assert.Equal(t, uint8(0xCA), c.A.Get())
	assert.Equal(t, uint8(0xFD), c.F.Get())

	c.BC.Set(0xCAFE)
	c.BC.Dec()
	assert.Equal(t, uint16(0xCAFD), c.BC.Get())
	assert.Equal(t, uint8(0xCA), c.B.Get())
	assert.Equal(t, uint8(0xFD), c.C.Get())

	c.DE.Set(0xCAFE)
	c.DE.Dec()
	assert.Equal(t, uint16(0xCAFD), c.DE.Get())
	assert.Equal(t, uint8(0xCA), c.D.Get())
	assert.Equal(t, uint8(0xFD), c.E.Get())

	c.HL.Set(0xCAFE)
	c.HL.Dec()
	assert.Equal(t, uint16(0xCAFD), c.HL.Get())
	assert.Equal(t, uint8(0xCA), c.H.Get())
	assert.Equal(t, uint8(0xFD), c.L.Get())

	c.SP.Set(0xFEFE)
	c.SP.Dec()
	assert.Equal(t, uint16(0xFEFD), c.SP.Get())

	c.PC.Set(0xFEFE)
	c.PC.Dec()
	assert.Equal(t, uint16(0xFEFD), c.PC.Get())
}

func TestWordRegistersCanDecreaseTheirValueWithOverflow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mem := mocks.NewMockMemoryInterface(ctrl)
	c := cpu.New(mem)

	c.AF.Set(0x0001)
	c.AF.DecBy(3)
	assert.Equal(t, uint16(0xFFFE), c.AF.Get())
	assert.Equal(t, uint8(0xFF), c.A.Get())
	assert.Equal(t, uint8(0xFE), c.F.Get())

	c.BC.Set(0x0001)
	c.BC.DecBy(3)
	assert.Equal(t, uint16(0xFFFE), c.BC.Get())
	assert.Equal(t, uint8(0xFF), c.B.Get())
	assert.Equal(t, uint8(0xFE), c.C.Get())

	c.DE.Set(0x0001)
	c.DE.DecBy(3)
	assert.Equal(t, uint16(0xFFFE), c.DE.Get())
	assert.Equal(t, uint8(0xFF), c.D.Get())
	assert.Equal(t, uint8(0xFE), c.E.Get())

	c.HL.Set(0x0001)
	c.HL.DecBy(3)
	assert.Equal(t, uint16(0xFFFE), c.HL.Get())
	assert.Equal(t, uint8(0xFF), c.H.Get())
	assert.Equal(t, uint8(0xFE), c.L.Get())

	c.SP.Set(0x0001)
	c.SP.DecBy(3)
	assert.Equal(t, uint16(0xFFFE), c.SP.Get())

	c.PC.Set(0x0001)
	c.PC.DecBy(3)
	assert.Equal(t, uint16(0xFFFE), c.PC.Get())
}
