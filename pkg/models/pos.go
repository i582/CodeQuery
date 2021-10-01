package models

import (
	"fmt"
)

type Pos struct {
	Filename string
	Line     int32
	EndLine  int32
	Col      int32
	Length   int32 // body length
}

func (p Pos) String() string {
	return fmt.Sprintf("%s:%d", p.Filename, p.Line)
}
