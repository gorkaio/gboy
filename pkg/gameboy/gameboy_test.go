package gameboy_test

import (
	"github.com/gorkaio/gboy/pkg/gameboy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialisesGameboySystem(t *testing.T) {
	_, err := gameboy.New()
	assert.NoError(t, err)
}
