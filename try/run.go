package try

// Run runs the specified function (f). If the function throws a panic, the handlers are tried in order to handle
// the panic. The handler functions can be custom, or one of the built-in error handlers, such as catch.ErrorByType,
// catch.ErrorByValue, catch.AnyError, or catch.Any. If no suitable panic handler is found the panic is re-thrown.
func Run(
	f func(),
	catch1 CatchHandler,
	catch ...CatchHandler,
) {
	result := run(f)
	if result != nil {
		for _, h := range append([]CatchHandler{catch1}, catch...) {
			if h(result) {
				return
			}
		}
		panic(result)
	}
}

// CatchHandler is a function that can process a panic value. The "catch" package offers a set of comprehensive
// handlers.
type CatchHandler func(e any) bool

func run(f func()) (result any) {
	defer func() {
		result = recover()
	}()
	f()
	return
}
