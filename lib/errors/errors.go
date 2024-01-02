package errors

import (
	"ConnectorSchema/lib/types"
	"fmt"

	"github.com/fatih/color"
)

type Error struct {
	Type    types.ErrorType
	Message string
}

func NewError(errType types.ErrorType, message string) *Error {
	return &Error{
		Type:    errType,
		Message: message,
	}
}

func (e *Error) Error() string {
	red := color.New(color.FgRed).SprintFunc()
	return fmt.Sprintf("%s: %s", red(e.Type.String()), e.Message)
}

func (e *Error) Print() {
	fmt.Printf("%s: %s\n", e.Type.String(), e.Message)
}

func (e *Error) Throw() {
	panic(e)
}
