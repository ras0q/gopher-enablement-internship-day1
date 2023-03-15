package a

import (
	"fmt"

	"go.uber.org/multierr" // want "hogehoge"
)

func f() {
	multierr.Errors(fmt.Errorf("hoge"))
}
