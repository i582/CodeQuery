package tester

import (
	"testing"

	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/internal/query"
	"github.com/i582/CodeQuery/pkg/querylang/token"
	"gotest.tools/assert"
)

type Lexer interface {
	Lex() *token.Token
}

type LexerTokenFreeFloatingTestSuite struct {
	t *testing.T

	Code     string
	Expected [][]*token.Token
}

func NewLexerTokenFreeFloatingTestSuite(t *testing.T) *LexerTokenFreeFloatingTestSuite {
	return &LexerTokenFreeFloatingTestSuite{
		t: t,
	}
}

func (l *LexerTokenFreeFloatingTestSuite) Run() {
	lexer := query.NewLexer([]byte(l.Code), conf.Config{})

	for _, expected := range l.Expected {
		tkn := lexer.Lex()
		actual := tkn.FreeFloating
		for _, v := range actual {
			v.Pos = nil
		}
		assert.DeepEqual(l.t, expected, actual)
	}
}
