package query

import (
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/position"
	"github.com/i582/CodeQuery/pkg/querylang/token"
)

type ParserBrackets struct {
	Pos             *position.Pos
	OpenBracketTkn  *token.Token
	Child           ast.Node
	CloseBracketTkn *token.Token
}

func (n *ParserBrackets) Walk(v ast.Visitor) {
	// do nothing
}

func (n *ParserBrackets) GetPos() *position.Pos {
	return n.Pos
}

type ParserSeparatedList struct {
	Pos           *position.Pos
	Items         []ast.Node
	SeparatorTkns []*token.Token
}

func (n *ParserSeparatedList) Walk(v ast.Visitor) {
	// do nothing
}

func (n *ParserSeparatedList) GetPos() *position.Pos {
	return n.Pos
}

// TraitAdaptationList node
type TraitAdaptationList struct {
	Pos                  *position.Pos
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []ast.Node
	CloseCurlyBracketTkn *token.Token
}

func (n *TraitAdaptationList) Walk(v ast.Visitor) {
	// do nothing
}

func (n *TraitAdaptationList) GetPos() *position.Pos {
	return n.Pos
}

// ArgumentList node
type ArgumentList struct {
	Pos                 *position.Pos
	OpenParenthesisTkn  *token.Token
	Arguments           []ast.Node
	SeparatorTkns       []*token.Token
	EllipsisTkn         *token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ArgumentList) Walk(v ast.Visitor) {
	// do nothing
}

func (n *ArgumentList) GetPos() *position.Pos {
	return n.Pos
}

type EnumCaseExpr struct {
	Pos       *position.Pos
	AssignTkn *token.Token
	Expr      ast.Node
}

func (n *EnumCaseExpr) Walk(v ast.Visitor) {
	// do nothing
}

func (n *EnumCaseExpr) GetPos() *position.Pos {
	return n.Pos
}

type ReturnType struct {
	Pos      *position.Pos
	ColonTkn *token.Token
	Type     ast.Node
}

func (n *ReturnType) Walk(v ast.Visitor) {
	// do nothing
}

func (n *ReturnType) GetPos() *position.Pos {
	return n.Pos
}

// TraitMethodRef node
type TraitMethodRef struct {
	Pos            *position.Pos
	Trait          ast.Node
	DoubleColonTkn *token.Token
	Method         ast.Node
}

func (n *TraitMethodRef) Walk(v ast.Visitor) {
	// do nothing
}

func (n *TraitMethodRef) GetPos() *position.Pos {
	return n.Pos
}
