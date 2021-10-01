package models

type IFuncArg interface {
	String() string
	IsInt() bool
	IsFloat() bool
	IsString() bool
	IsBool() bool
	IsConstant() bool
	IsVariable() bool
	IsExpression() bool
}

type IFuncCall interface {
	Arg(index int64) IFuncArg
	Args_() int64
	File() string
}

type IFuncPath interface {
	Begin() IFunc
	End() IFunc
	At(index int64) IFunc
	Length() int64
	Contains(calls IFuncTable) bool
}
