package exceptions_test

import (
	"github.com/apitalist/exceptions"
	"github.com/apitalist/exceptions/try"
	"github.com/apitalist/exceptions/try/catch"
	"strings"
	"testing"
)

func throwtest1() {
	throwtest2()
}

func throwtest2() {
	panic(exceptions.New("Hello world!"))
}

func TestE2E(t *testing.T) {
	caught := false
	try.Run(
		throwtest1,
		catch.Exception(func(e exceptions.Exception) {
			stack := e.Stack()
			i := 0
			caught = true
			for {
				frame, next := stack.Next()
				switch i {
				case 0:
					if !strings.HasSuffix(frame.Function, ".throwtest2") {
						t.Fatalf("Incorrect first stack item: %s", frame.Function)
					}
				case 1:
					if !strings.HasSuffix(frame.Function, ".throwtest1") {
						t.Fatalf("Incorrect second stack item: %s", frame.Function)
					}
				}
				if !next {
					break
				}
				i++
			}
		}))
	if !caught {
		t.Fail()
	}
}
