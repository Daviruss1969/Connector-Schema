package errors

import (
	"PAValidator/lib/types"
	"fmt"

	"github.com/fatih/color"
)

type Error struct {
	Type    types.ErrorType
	Line    int
	Message string
}

func NewError(errType types.ErrorType, message string) *Error {
	return &Error{
		Type:    errType,
		Message: message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type.String(), e.Message)
}

func (e *Error) Print() {
	fmt.Printf("%s: %s\n", color.New(color.FgRed, color.Bold).Sprintf("%s", e.Type.String()), e.Message)
}

func (e *Error) Throw() {
	panic(e)
}
