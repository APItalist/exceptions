package catch

import (
	"github.com/apitalist/exceptions/try"
)

// Any creates a catch-all error handler.
func Any(f func(err any)) try.CatchHandler {
	return func(e any) bool {
		f(e)
		return true
	}
}
