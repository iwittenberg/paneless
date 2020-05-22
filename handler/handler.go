package handler

import (
	"errors"
	a "github.com/iwittenberg/paneless/arrangements"
)

// Handler is an interface to any ability to rearrange windows.
type Handler interface {
	Apply(a *a.Arrangement)
	GetCurrentWindowPositions() *a.Arrangement
	OpenFile(file string)
	RegisterHotkeys(as *a.Arrangements)
}

type OS int

const (
	WINDOWS = iota
)

// NewHandler takes an OS and returns an implementation of the Handler interface for the appropriate operating system.
func NewHandler(os OS) (Handler, error) {
	switch os {
	case WINDOWS:
		return Windows{}, nil
	default:
		return nil, errors.New("invalid OS type when creating new Handler")
	}
}
