package models

import (
	"github.com/i582/CodeQuery/pkg/querylang/models"
)

type FuncPath []*Func

func (f FuncPath) Begin() models.IFunc {
	return f[0]
}

func (f FuncPath) End() models.IFunc {
	return f[len(f)-1]
}

func (f FuncPath) At(index int64) models.IFunc {
	return f[index]
}

func (f FuncPath) Length() int64 {
	return int64(len(f))
}

func (f FuncPath) Contains(funcsRaw models.IFuncTable) bool {
	funcs, ok := funcsRaw.(*FuncTable)
	if !ok {
		return false
	}
	funcMap := make(map[int64]struct{}, len(funcs.Data))
	for _, fun := range funcs.Data {
		funcMap[fun.ID] = struct{}{}
	}

	for _, fun := range f {
		if _, ok := funcMap[fun.ID]; ok {
			return true
		}
	}

	return false
}
