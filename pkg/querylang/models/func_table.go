package models

type IFuncTable interface {
	Iterate(func(fun IFunc))
	Length() int64
}
