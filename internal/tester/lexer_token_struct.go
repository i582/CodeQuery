package tester

import (
	"testing"

	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/internal/query"
	"github.com/i582/CodeQuery/pkg/querylang/token"
	"gotest.tools/assert"
)

type LexerTokenStructTestSuite struct {
	t *testing.T

	Code     string
	Expected []*token.Token

	withPos          bool
	withFreeFloating bool
}

func NewLexerTokenStructTestSuite(t *testing.T) *LexerTokenStructTestSuite {
	return &LexerTokenStructTestSuite{
		t: t,
	}
}

func (l *LexerTokenStructTestSuite) WithPos() {
	l.withPos = true
}

func (l *LexerTokenStructTestSuite) WithFreeFloating() {
	l.withFreeFloating = true
}

func (l *LexerTokenStructTestSuite) Run() {
	lexer := query.NewLexer([]byte(l.Code), conf.Config{})

	for _, expected := range l.Expected {
		actual := lexer.Lex()
		if !l.withPos {
			actual.Pos = nil
		}
		if !l.withFreeFloating {
			actual.FreeFloating = nil
		}
		assert.DeepEqual(l.t, expected, actual)
	}
}
