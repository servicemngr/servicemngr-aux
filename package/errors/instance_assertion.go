package rherrors

import (
	"errors"
	"github.com/servicemngr/core/package/instance"
)

func NewInstanceAssertionError(i interface{ ID() instance.ID }) error {
	return errors.New("Failed to assert instance " + i.ID().String())
}
