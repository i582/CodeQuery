package token

import (
	"github.com/i582/CodeQuery/pkg/querylang/position"
)

//go:generate stringer -type=ID -output ./token_string.go
type ID int

const (
	T_SELECT ID = iota + 57346
	T_FROM
	T_WHERE
	T_WITH
	T_COUNT
	T_LIMIT
	T_OFFSET
	T_ORDER_BY
	T_DESC
	T_ASC

	T_LNUMBER
	T_DNUMBER
	T_CONSTANT_STRING
	T_STRING
	T_VARIABLE
	T_COMMENT
	T_DOC_COMMENT
	T_WHITESPACE
	T_OR
	T_XOR
	T_AND
	T_NOT
	T_EQUAL
	T_NOT_EQUAL
	T_SMALLER
	T_GREATER
	T_SMALLER_OR_EQUAL
	T_GREATER_OR_EQUAL
	T_OBJECT_OPERATOR
)

type Token struct {
	ID           ID
	Value        []byte
	Pos          *position.Pos
	FreeFloating []*Token
}

func (t *Token) GetPos() *position.Pos {
	return t.Pos
}
