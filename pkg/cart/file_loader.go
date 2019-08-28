package cart

import "io/ioutil"

// Loader defines the interface for a cart loader
type Loader interface {
	Load(resource string) (CartInterface, error)
}

// FileLoader struct
type FileLoader struct {
	romfile string
}

// NewFileLoader initialises a new romfile loader
func NewFileLoader() Loader {
	return &FileLoader{}
}

// Load loads a romfile from file
func (fl *FileLoader) Load(resource string) (CartInterface, error) {
	data, err := ioutil.ReadFile(resource)
	if err != nil {
		return nil, err
	}

	cart, err := newCart(data)
	if err != nil {
		return nil, err
	}

	fl.romfile = resource
	return cart, nil
}
