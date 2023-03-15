package a

import (
	"go.uber.org/multierr"
)

func f() {
	multierr.Errors() // want "error found"
}
