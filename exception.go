package exceptions

import (
	"runtime"
)

const DefaultExceptionStackDepth = 100

// New creates a new Exception from a message provided.
func New(message string) Exception {
	return &exception{
		Inherit(nil, 1, DefaultExceptionStackDepth),
		message,
	}
}

func NewFrom(cause error) Exception {
	return &exception{
		Inherit(cause, 1, DefaultExceptionStackDepth),
		cause.Error(),
	}
}

func NewFromWithMessage(message string, cause error) Exception {
	return &exception{
		Inherit(cause, 1, DefaultExceptionStackDepth),
		message,
	}
}

// BaseException contains a stack trace and potentially a cause (original error) for the exception built from this base.
// However, the BaseException doesn't satisfy the Exception interface and must be enriched with an Error() function
// to serve as an exception. You can use the Inherit() function to create a BaseException suitable for embedding in a
// struct.
type BaseException interface {
	// Stack returns the call stack from the point where the exception was created.
	Stack() *runtime.Frames
	// Unwrap returns the original error that caused this exception.
	Unwrap() error
}

// Exception is a combination of the stack trace and unwrapping functionality contained in BaseException and the
// traditional Go error interface. You can use this like you would use a traditional Go error, or in a panic()
// with the handling facilities of this library.
type Exception interface {
	error

	BaseException
}

type exception struct {
	BaseException

	message string
}

func (e exception) Error() string {
	return e.message
}

// Inherit is a helper function to create an exception that you can embed into other
// exceptions as a method of inheritance.
func Inherit(cause error, skipStackItems int, stackItems int) BaseException {
	if stackItems < 0 {
		stackItems = DefaultExceptionStackDepth
	}
	rpc := make([]uintptr, stackItems)
	runtime.Callers(skipStackItems + 2, rpc[:])
	return &baseException{
		cause,
		rpc,
	}
}

type baseException struct {
	cause error
	pcs []uintptr
}

func (e baseException) Unwrap() error {
	return e.cause
}

func (e baseException) Stack() *runtime.Frames {
	return runtime.CallersFrames(e.pcs)
}

