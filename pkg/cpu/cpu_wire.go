//+build wireinject

package cpu

import (
	"github.com/google/wire"
	"github.com/gorkaio/gboy/pkg/memory"
)

func InitializeCPU() (CPUInterface, error) {
	wire.Build(New, memory.InitializeMemory)
	return nil, nil
}
