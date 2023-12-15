package catch

import (
	"github.com/apitalist/exceptions/try"
)

// ByType creates a panic handler that handles only one specific type of item throw in the panic. This is useful when
// you want to panic with non-error types, such as:
//
//	type myCustomType struct {
//	}
func ByType[T any](f func(err T)) try.CatchHandler {
	return func(e any) bool {
		err, ok := e.(T)
		if !ok {
			return false
		}
		f(err)
		return true
	}
}
