package gameboy_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/gorkaio/gboy/pkg/gameboy/mocks"
	"github.com/gorkaio/gboy/pkg/gameboy"
	"github.com/stretchr/testify/assert"
)

func TestInitialisesGameboySystem(t *testing.T) {
	ctrlMemory := gomock.NewController(t)
	defer ctrlMemory.Finish()
	memory := mocks.NewMockMemory(ctrlMemory)

	ctrlCPU := gomock.NewController(t)
	defer ctrlCPU.Finish()
	cpu := mocks.NewMockCPU(ctrlCPU)

	_, err := gameboy.New(memory, cpu)
	assert.NoError(t, err)
}
