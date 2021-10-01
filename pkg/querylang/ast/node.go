package ast

type Node interface {
	Walk(v Visitor)
}

type Visitor interface {
	EnterNode(n Node) bool
	LeaveNode(n Node)
}
