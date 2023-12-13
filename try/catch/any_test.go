package catch_test

import (
	"fmt"

	"github.com/apitalist/exceptions/try"
	"github.com/apitalist/exceptions/try/catch"
)

func ExampleAny() {
	try.Run(
		func() {
			// Panic with an error:
			panic("something bad happened")
		},
		// Handle a specific error type:
		catch.Any(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught panic: something bad happened
}
