//+build wireinject

package cart

import (
	"github.com/google/wire"
)

func InitializeLoader() (Loader, error) {
	wire.Build(NewFileLoader)
	return nil, nil
}
