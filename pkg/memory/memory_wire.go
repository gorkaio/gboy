//+build !wireinject

package memory

import (
	"github.com/google/wire"
)

// InitializeMemory initializes a new memory
func InitializeMemory() (*Memory, error) {
	wire.Build(New)
	return nil, nil
}
