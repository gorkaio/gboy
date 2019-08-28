//+build wireinject

package memory

import (
	"github.com/gorkaio/gboy/pkg/cart"
	"github.com/google/wire"
)

func InitializeMemory() (MemoryInterface, error) {
	wire.Build(New, cart.NewFileLoader)
	return nil, nil
}
