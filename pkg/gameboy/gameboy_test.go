package gameboy_test

import (
	gameboy "github.com/gorkaio/gboy/pkg/gameboy"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

const testfile string = "../../roms/10-print.gb"

func TestInitialisesGameboySystem(t *testing.T) {
	_, err := gameboy.New(testfile)
	assert.NoError(t, err)
}
