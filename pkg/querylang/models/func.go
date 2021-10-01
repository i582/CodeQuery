package models

type IFuncCalls interface {
	Count() int64
}

type IFQN interface {
	ClassName() string
	Name() string
	Namespace() string
	FQN() string
}

type IFunc interface {
	Name() string
	Namespace() string
	ClassName() string
	FullName() string

	Globals() IGlobals
	Calls() IFuncCalls
	CountUse() int64
}
