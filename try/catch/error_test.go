package catch_test

import (
	"fmt"

	"github.com/apitalist/exceptions/try"
	"github.com/apitalist/exceptions/try/catch"
)

func ExampleError() {
	try.Run(
		func() {
			// Panic with an error:
			panic(fmt.Errorf("something bad happened"))
		},
		// Handle a specific error type:
		catch.Error(
			func(err error) {
				fmt.Printf("Caught error: %v", err)
			},
		),
	)

	// Output: Caught error: something bad happened
}
