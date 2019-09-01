//+build !wireinject

package gameboy

import (
	"github.com/gorkaio/gboy/pkg/cpu"
	"github.com/gorkaio/gboy/pkg/memory"
	"github.com/google/wire"
)

func InitializeGameboy() (*Gameboy, error) {
	wire.Build(New, memory.InitializeMemory, cpu.InitializeCPU)
	return nil, nil
}
