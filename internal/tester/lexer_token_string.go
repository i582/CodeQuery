package tester

import (
	"testing"

	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/internal/query"
	"gotest.tools/assert"
)

type LexerTokenStringTestSuite struct {
	t *testing.T

	Code     string
	Expected []string
}

func NewLexerTokenStringTestSuite(t *testing.T) *LexerTokenStringTestSuite {
	return &LexerTokenStringTestSuite{
		t: t,
	}
}

func (l *LexerTokenStringTestSuite) Run() {
	lexer := query.NewLexer([]byte(l.Code), conf.Config{})

	for _, expected := range l.Expected {
		tkn := lexer.Lex()
		actual := string(tkn.Value)
		assert.DeepEqual(l.t, expected, actual)
	}
}
