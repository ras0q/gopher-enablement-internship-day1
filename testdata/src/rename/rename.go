package rename

import (
	"fmt"

	me "go.uber.org/multierr"
)

func rename() {
	me.Errors(fmt.Errorf("hoge")) // want "multierr found"
}
