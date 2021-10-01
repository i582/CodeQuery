package ast

import (
	"github.com/i582/CodeQuery/pkg/querylang/position"
	"github.com/i582/CodeQuery/pkg/querylang/token"
)

type (
	Root struct {
		Pos   *position.Pos
		Stmts []Node
	}

	CommaSeparatedList struct {
		Pos     *position.Pos
		Items   []Node
		SepTkns []*token.Token
	}

	Identifier struct {
		Pos           *position.Pos
		IdentifierTkn *token.Token
		Value         []byte
	}

	Variable struct {
		Pos     *position.Pos
		NameTkn *token.Token
		Name    []byte
	}

	BasicLit struct {
		Pos   *position.Pos
		Kind  *token.Token
		Value []byte
	}

	ComparisonExpr struct {
		Pos   *position.Pos
		Left  Node
		OpTkn *token.Token
		Right Node
	}

	BinaryExpr struct {
		Pos   *position.Pos
		Left  Node
		OpTkn *token.Token
		Right Node
	}

	NotExpr struct {
		Pos    *position.Pos
		NotTkn *token.Token
		Expr   Node
	}

	MethodCallExpr struct {
		Pos             *position.Pos
		Variable        Node
		OpTkn           *token.Token
		MethodName      *Identifier
		OpenBracketTkn  *token.Token
		Args            []Node
		SeparatorsTkns  []*token.Token
		CloseBracketTkn *token.Token
	}

	SelectExpr struct {
		Pos         *position.Pos
		SelectTkn   *token.Token
		Select      *SelectSubjectExpr
		FromExpr    *FromExpr
		WhereExpr   *WhereExpr
		WithExpr    *WithExpr
		LimitExpr   *LimitExpr
		OrderByExpr *OrderByExpr
	}

	SelectSubjectExpr struct {
		Pos   *position.Pos
		List  *CommaSeparatedList // <Identifier>
		Star  *token.Token        // if SELECT *
		Count *token.Token        // if SELECT COUNT(*)
	}

	FromExpr struct {
		Pos     *position.Pos
		FromTkn *token.Token
		From    Node
	}

	WhereExpr struct {
		Pos      *position.Pos
		WhereTkn *token.Token
		Expr     Node
	}

	WithExpr struct {
		Pos      *position.Pos
		WithTkn  *token.Token
		WithList *CommaSeparatedList // <Expr>
	}

	LimitExpr struct {
		Pos      *position.Pos
		LimitTkn *token.Token
		Value    Node
	}

	OrderByExpr struct {
		Pos        *position.Pos
		OrderByTkn *token.Token
		Field      Node
		Desc       bool
		Asc        bool
	}
)
