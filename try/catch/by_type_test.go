package catch_test

import (
	"fmt"
	"github.com/apitalist/exceptions/try"
	"github.com/apitalist/exceptions/try/catch"
)

type myCustomPanicType struct {
	num int
}

func ExampleByType() {
	try.Run(
		func() {
			panic(myCustomPanicType{42})
		},
		// Run the specific data type, which may or may not be an error:
		catch.ByType(func(data myCustomPanicType) {
			fmt.Println(data.num)
		}),
	)
	// Output: 42
}
