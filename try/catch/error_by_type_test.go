package catch_test

import (
	"fmt"

	"github.com/apitalist/exceptions/try"
	"github.com/apitalist/exceptions/try/catch"
)

type myCustomErrorType struct {
}

func (m myCustomErrorType) Error() string {
	return "My custom error"
}

func ExampleErrorByType() {
	try.Run(
		func() {
			// Panic with an error:
			var err error = &myCustomErrorType{}
			panic(err)
		},
		// Handle a specific error type:
		catch.ErrorByType(
			func(err *myCustomErrorType) {
				fmt.Printf("Caught my custom error: %v", err)
			},
		),
	)

	// Output: Caught my custom error: My custom error
}
