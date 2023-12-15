package exceptions

import (
	"fmt"
	"runtime"
)

const ExceptionStackDepth = 100

func New(message string, args ...any) Exception {
	return NewBaseException(1, message, args...)
}

type Exception interface {
	error

	Stack() *runtime.Frames
}

// NewBaseException is the constructor for creating an inherited exception.
//
//     type MyException struct {
//         *exceptions.BaseException
//     }
//
//     func NewMyException() Exception {
//         return &MyException{
//             exceptions.NewBaseException(
//                 // How many stack items to skip:
//                 1,
//                 // Message:
//                 "Hello world!"
//             ),
//         }
//     }
func NewBaseException(skipStackItems int, message string, args ...any) *BaseException {
	rpc := make([]uintptr, ExceptionStackDepth)
	runtime.Callers(skipStackItems + 2, rpc[:])
	return &BaseException{
		fmt.Errorf(message, args...),
		rpc,
	}
}

type BaseException struct {
	err error
	pcs []uintptr
}

func (e BaseException) Error() string {
	return e.err.Error()
}

func (e BaseException) Stack() *runtime.Frames {
	return runtime.CallersFrames(e.pcs)
}

