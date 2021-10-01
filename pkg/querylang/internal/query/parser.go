package query

import (
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/errors"
	"github.com/i582/CodeQuery/pkg/querylang/internal/position"
	"github.com/i582/CodeQuery/pkg/querylang/token"
)

// Parser structure
type Parser struct {
	Lexer          *Lexer
	currentToken   *token.Token
	rootNode       ast.Node
	errHandlerFunc func(*errors.Error)
	builder        *Builder
}

// NewParser creates and returns new Parser
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		Lexer:          lexer,
		errHandlerFunc: lexer.errHandlerFunc,
	}
	p.builder = NewBuilder(position.NewBuilder(), p)
	return p
}

func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex()

	p.currentToken = t
	lval.token = t

	return int(t.ID)
}

func (p *Parser) Error(msg string) {
	if p.errHandlerFunc == nil {
		return
	}

	p.errHandlerFunc(errors.NewError(msg, p.currentToken.Pos))
}

// Parse the php7 Parser entrypoint
func (p *Parser) Parse() int {
	p.rootNode = nil

	return yyParse(p)
}

// GetRootNode returns root node
func (p *Parser) GetRootNode() ast.Node {
	return p.rootNode
}

// helpers

func lastNode(nn []ast.Node) ast.Node {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}
