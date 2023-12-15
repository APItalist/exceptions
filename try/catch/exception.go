package catch

import (
	"errors"
	"github.com/apitalist/exceptions"
	"github.com/apitalist/exceptions/try"
)

func Exception[T exceptions.Exception](f func(e T)) try.CatchHandler {
	return func(e any) bool {
		err, ok := e.(error)
		if !ok {
			return false
		}
		var errorType T
		if errors.As(err, &errorType) {
			f(errorType)
			return true
		}
		return false
	}
}
