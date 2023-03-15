package a

import (
	"fmt"

	"go.uber.org/multierr"
)

func f() {
	multierr.Errors(fmt.Errorf("hoge")) // want "multierr found"
}
