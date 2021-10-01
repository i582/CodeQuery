package models

import (
	"github.com/i582/CodeQuery/pkg/querylang/models"
)

type TableCol struct {
	Name   string
	IdName string
	Hide   bool
}

type FuncTable struct {
	Cols []TableCol
	Data []*Func
}

func (t *FuncTable) Iterate(f func(fun models.IFunc)) {
	for _, d := range t.Data {
		f(d)
	}
}

func (t *FuncTable) Length() int64 {
	return int64(len(t.Data))
}

func (t *FuncTable) NeedShowCol(name string) bool {
	for _, col := range t.Cols {
		if col.IdName == name {
			return true
		}
	}
	return false
}

type FuncMap map[FQN]*Func

type FuncCallArgType uint8

const (
	IntArg FuncCallArgType = iota
	FloatArg
	BoolArg
	StringArg
	VariableArg
	ConstantArg
	ExpressionArg
)

type FuncCallArg struct {
	Type  FuncCallArgType
	Value string
}

func (a FuncCallArg) String() string     { return a.Value }
func (a FuncCallArg) IsInt() bool        { return a.Type == IntArg }
func (a FuncCallArg) IsFloat() bool      { return a.Type == FloatArg }
func (a FuncCallArg) IsBool() bool       { return a.Type == BoolArg }
func (a FuncCallArg) IsString() bool     { return a.Type == StringArg }
func (a FuncCallArg) IsVariable() bool   { return a.Type == VariableArg }
func (a FuncCallArg) IsConstant() bool   { return a.Type == ConstantArg }
func (a FuncCallArg) IsExpression() bool { return a.Type == ExpressionArg }

type FuncCallArgs []FuncCallArg

type FuncCall struct {
	ID int64

	Pos   Pos
	Args_ FuncCallArgs
}

func (c FuncCall) Arg(index int64) models.IFuncArg {
	return c.Args_[index]
}

func (c FuncCall) Args() int64 {
	return int64(len(c.Args_))
}

func (c FuncCall) File() string {
	return c.Pos.String()
}

type FuncCallTable struct {
	Cols []TableCol
	Data []FuncCall
}

func (t *FuncCallTable) NeedShowCol(name string) bool {
	for _, col := range t.Cols {
		if col.IdName == name {
			return true
		}
	}
	return false
}

type FuncCalls []FuncCall

func (f FuncCalls) Count() int64 {
	return int64(len(f))
}

type Func struct {
	ID int64

	Fqn  FQN
	Pos_ Pos
	Rem  bool

	Globals_ Globals

	Calls_ FuncCalls

	UseCount int64
}

func (f Func) FullName() string {
	return f.Fqn.FQN()
}

func (f Func) Name() string {
	return f.Fqn.Name()
}

func (f Func) Namespace() string {
	return f.Fqn.Namespace()
}

func (f Func) ClassName() string {
	return f.Fqn.ClassName()
}

func (f Func) Pos() Pos {
	return f.Pos_
}

func (f Func) Globals() models.IGlobals {
	return f.Globals_
}

func (f Func) Calls() models.IFuncCalls {
	return f.Calls_
}

func (f Func) CountUse() int64 {
	return f.UseCount
}
