package models

type IGlobal interface {
	Id() int64
	Name() string
	CountUse() int64
}

type IGlobals interface {
	Count() int64
	Has(global IGlobal) bool
	Contains(globals IGlobals) bool
	Iterate(func(global IGlobal) bool)
}
