package query

import (
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/internal/position"
	"github.com/i582/CodeQuery/pkg/querylang/token"
)

// Builder is responsible for creating nodes inside grammar rules.
// Creating nodes directly in grammar rules is inconvenient, since
// there is no autocompletion, and you cannot put breakpoints inside.
type Builder struct {
	Pos    *position.Builder
	Parser *Parser
}

// NewBuilder creates a new Builder.
func NewBuilder(pos *position.Builder, parser *Parser) *Builder {
	return &Builder{
		Pos:    pos,
		Parser: parser,
	}
}

// NewEmptySeparatedList creates a new empty list.
// Used for places where a delimited list is used.
func (b *Builder) NewEmptySeparatedList() *ast.CommaSeparatedList {
	return &ast.CommaSeparatedList{}
}

// NewSeparatedList creates a new single-element list.
// Used for places where a delimited list is used.
func (b *Builder) NewSeparatedList(node ast.Node) *ast.CommaSeparatedList {
	return &ast.CommaSeparatedList{Items: []ast.Node{node}}
}

// NewSeparatedListWithTwoElements creates a new two-element list.
// Used for places where a delimited list is used.
func (b *Builder) NewSeparatedListWithTwoElements(node1 ast.Node, tkn *token.Token, node2 ast.Node) *ast.CommaSeparatedList {
	return &ast.CommaSeparatedList{
		Items:   []ast.Node{node1, node2},
		SepTkns: []*token.Token{tkn},
	}
}

func (b *Builder) NewNonEmptySeparatedList(nodes []ast.Node, tkns []*token.Token) *ast.CommaSeparatedList {
	return &ast.CommaSeparatedList{
		Items:   nodes,
		SepTkns: tkns,
	}
}

// AppendToSeparatedList inserts a new node and/or token into the list.
func (b *Builder) AppendToSeparatedList(list ast.Node, tkn *token.Token, node ast.Node) *ast.CommaSeparatedList {
	sepList := list.(*ast.CommaSeparatedList)
	if node != nil {
		sepList.Items = append(sepList.Items, node)
	}
	if tkn != nil {
		sepList.SepTkns = append(sepList.SepTkns, tkn)
	}

	return sepList
}

func (b *Builder) SeparatedListItems(list ast.Node) (items []ast.Node, sepTkns []*token.Token) {
	if list == nil {
		return nil, nil
	}

	paramsList, ok := list.(*ast.CommaSeparatedList)
	if !ok {
		return nil, nil
	}

	return paramsList.Items, paramsList.SepTkns
}

func (b *Builder) NewRoot(
	Stmts []ast.Node,
) *ast.Root {
	return &ast.Root{
		Pos:   nil,
		Stmts: Stmts,
	}
}

func (b *Builder) NewLimitExpr(
	LimitTkn *token.Token,
	Value ast.Node,
) *ast.LimitExpr {
	return &ast.LimitExpr{
		Pos:      nil,
		LimitTkn: LimitTkn,
		Value:    Value,
	}
}

func (b *Builder) NewOrderByExpr(
	OrderByTkn *token.Token,
	Value ast.Node,
	DescOrAscTkn *token.Token,
) *ast.OrderByExpr {
	asc := true
	if DescOrAscTkn != nil {
		asc = DescOrAscTkn.ID == token.T_ASC
	}
	return &ast.OrderByExpr{
		Pos:        nil,
		OrderByTkn: OrderByTkn,
		Field:      Value,
		Asc:        asc,
		Desc:       !asc,
	}
}

func (b *Builder) NewFromExpr(
	FromTkn *token.Token,
	From ast.Node,
) *ast.FromExpr {
	return &ast.FromExpr{
		Pos:     nil,
		FromTkn: FromTkn,
		From:    From,
	}
}

func (b *Builder) NewWhereExpr(
	WhereTkn *token.Token,
	Expr ast.Node,
) *ast.WhereExpr {
	return &ast.WhereExpr{
		Pos:      nil,
		WhereTkn: WhereTkn,
		Expr:     Expr,
	}
}

func (b *Builder) NewWithExpr(
	WithTkn *token.Token,
	WithList ast.Node,
) *ast.WithExpr {
	var list *ast.CommaSeparatedList
	if WithList != nil {
		list = WithList.(*ast.CommaSeparatedList)
	}
	return &ast.WithExpr{
		Pos:      nil,
		WithTkn:  WithTkn,
		WithList: list,
	}
}

func (b *Builder) NewSelectSubjectExpr(
	List ast.Node,
	Count *token.Token,
	Star *token.Token,
) *ast.SelectSubjectExpr {
	var list *ast.CommaSeparatedList
	if List != nil {
		list = List.(*ast.CommaSeparatedList)
	}

	return &ast.SelectSubjectExpr{
		Pos:   nil,
		List:  list,
		Star:  Star,
		Count: Count,
	}
}

func (b *Builder) NewSelectExpr(
	SelectTkn *token.Token,
	Select ast.Node,
	FromExpr ast.Node,
	WhereExpr ast.Node,
	WithExpr ast.Node,
	LimitExpr ast.Node,
	OrderByExpr ast.Node,
) *ast.SelectExpr {
	var from *ast.FromExpr
	var where *ast.WhereExpr
	var with *ast.WithExpr
	var limit *ast.LimitExpr
	var order *ast.OrderByExpr

	if FromExpr != nil {
		from = FromExpr.(*ast.FromExpr)
	}
	if WhereExpr != nil {
		where = WhereExpr.(*ast.WhereExpr)
	}
	if WithExpr != nil {
		with = WithExpr.(*ast.WithExpr)
	}
	if LimitExpr != nil {
		limit = LimitExpr.(*ast.LimitExpr)
	}
	if OrderByExpr != nil {
		order = OrderByExpr.(*ast.OrderByExpr)
	}

	return &ast.SelectExpr{
		SelectTkn:   SelectTkn,
		Select:      Select.(*ast.SelectSubjectExpr),
		FromExpr:    from,
		WhereExpr:   where,
		WithExpr:    with,
		LimitExpr:   limit,
		OrderByExpr: order,
	}
}

func (b *Builder) NewComparisonExpr(
	Variable ast.Node,
	OpTkn *token.Token,
	Expr ast.Node,
) *ast.ComparisonExpr {
	return &ast.ComparisonExpr{
		Pos:   nil,
		Left:  Variable,
		OpTkn: OpTkn,
		Right: Expr,
	}
}

func (b *Builder) NewBinaryExpr(
	Left ast.Node,
	OpTkn *token.Token,
	Right ast.Node,
) *ast.BinaryExpr {
	return &ast.BinaryExpr{
		Pos:   nil,
		Left:  Left,
		OpTkn: OpTkn,
		Right: Right,
	}
}

func (b *Builder) NewNotExpr(
	NotTkn *token.Token,
	Expr ast.Node,
) *ast.NotExpr {
	return &ast.NotExpr{
		Pos:    nil,
		NotTkn: NotTkn,
		Expr:   Expr,
	}
}

func (b *Builder) NewMethodCallExpr(
	Variable ast.Node,
	OpTkn *token.Token,
	MethodName *token.Token,
	OpenBracketTkn *token.Token,
	Args ast.Node,
	CloseBracketTkn *token.Token,
) *ast.MethodCallExpr {
	args, seps := b.SeparatedListItems(Args)
	return &ast.MethodCallExpr{
		Pos:             nil,
		Variable:        Variable,
		OpTkn:           OpTkn,
		MethodName:      b.NewIdentifier(MethodName),
		OpenBracketTkn:  OpenBracketTkn,
		Args:            args,
		SeparatorsTkns:  seps,
		CloseBracketTkn: CloseBracketTkn,
	}
}

func (b *Builder) NewBasicLit(
	LiteralTkn *token.Token,
) *ast.BasicLit {
	return &ast.BasicLit{
		Pos:   nil,
		Kind:  LiteralTkn,
		Value: LiteralTkn.Value,
	}
}

//
// func (b *Builder) NewExpressionStmt(
// 	Right ast.Node,
// 	SemiColonTkn *token.Token,
// ) ast.Node {
// 	return &ast.StmtExpression{
// 		Pos_:     b.Pos_.NewNodeTokenPos(Right, SemiColonTkn),
// 		Right:         Right,
// 		SemiColonTkn: SemiColonTkn,
// 	}
// }
//

func (b *Builder) NewVariable(
	VariableTkn *token.Token,
) *ast.Variable {
	if VariableTkn == nil {
		return nil
	}

	return &ast.Variable{
		Pos:     b.Pos.NewTokenPos(VariableTkn),
		NameTkn: VariableTkn,
		Name:    VariableTkn.Value,
	}
}

func (b *Builder) NewIdentifier(
	IdentifierTkn *token.Token,
) *ast.Identifier {
	if IdentifierTkn == nil {
		return nil
	}

	return &ast.Identifier{
		Pos:           b.Pos.NewTokenPos(IdentifierTkn),
		IdentifierTkn: IdentifierTkn,
		Value:         IdentifierTkn.Value,
	}
}

//
// func (b *Builder) NewMethodCall(
// 	Right ast.Node,
// 	ObjectOperatorTkn *token.Token,
// 	PropertyName ast.Node,
// 	ArgList ast.Node,
// ) *ast.ExprMethodCall {
// 	argumentList := ArgList.(*ArgumentList)
// 	methodCall := &ast.ExprMethodCall{
// 		Pos_:            b.Pos_.NewNodesPos(Right, ArgList),
// 		Var:                 Right,
// 		ObjectOperatorTkn:   ObjectOperatorTkn,
// 		Method:              PropertyName,
// 		OpenParenthesisTkn:  argumentList.OpenParenthesisTkn,
// 		Args:                argumentList.Arguments,
// 		SeparatorTkns:       argumentList.SeparatorTkns,
// 		EllipsisTkn:         argumentList.EllipsisTkn,
// 		CloseParenthesisTkn: argumentList.CloseParenthesisTkn,
// 	}
//
// 	if brackets, ok := PropertyName.(*ParserBrackets); ok {
// 		methodCall.OpenCurlyBracketTkn = brackets.OpenBracketTkn
// 		methodCall.Method = brackets.Child
// 		methodCall.CloseCurlyBracketTkn = brackets.CloseBracketTkn
// 	}
//
// 	return methodCall
// }
