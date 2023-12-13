package try_test

import (
	"fmt"
	"github.com/apitalist/exceptions/try"
	"github.com/apitalist/exceptions/try/catch"
)

func Example() {
	try.Run(
		func() {
			fmt.Println("Hello world!")
		},
		catch.Any(
			func(err any) {
				fmt.Printf("An error happened: %v", err)
			},
		),
	)
	// Output: Hello world!
}
