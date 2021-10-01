package errors

import (
	"fmt"

	"github.com/i582/CodeQuery/pkg/querylang/position"
)

// Error parsing error
type Error struct {
	Msg string
	Pos *position.Pos
}

// NewError creates and returns new Error
func NewError(msg string, p *position.Pos) *Error {
	return &Error{
		Msg: msg,
		Pos: p,
	}
}

func (e *Error) String() string {
	atLine := ""
	if e.Pos != nil {
		atLine = fmt.Sprintf(" at line %d", e.Pos.StartLine)
	}

	return fmt.Sprintf("%s%s", e.Msg, atLine)
}
