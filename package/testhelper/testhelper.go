package testhelper

import (
	"github.com/servicemngr/core/package/instance"
	"testing"
)

func GetErrorFunc(t *testing.T) instance.ErrorFunc {
	return func(e ...interface{}) {
		t.Error(e...)
	}
}
