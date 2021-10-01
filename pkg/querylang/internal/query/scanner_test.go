package query_test

import (
	"testing"

	"github.com/i582/CodeQuery/internal/tester"
	"github.com/i582/CodeQuery/pkg/querylang/token"
)

func TestSelect(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.Code = "SELECT * FROM"
	suite.Expected = []*token.Token{
		{
			ID:    token.T_SELECT,
			Value: []byte("SELECT"),
		},
		{
			ID:    '*',
			Value: []byte("*"),
		},
		{
			ID:    token.T_FROM,
			Value: []byte("FROM"),
		},
	}
	suite.Run()
}

func TestAssign(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.Code = `name="value"`
	suite.Expected = []*token.Token{
		{
			ID:    token.T_STRING,
			Value: []byte("name"),
		},
		{
			ID:    '=',
			Value: []byte("="),
		},
		{
			ID:    token.T_CONSTANT_STRING,
			Value: []byte(`"value"`),
		},
	}
	suite.Run()
}

func TestCountKeyword(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.Code = "calls.count()"
	suite.Expected = []*token.Token{
		{
			ID:    token.T_STRING,
			Value: []byte("calls"),
		},
		{
			ID:    token.T_OBJECT_OPERATOR,
			Value: []byte("."),
		},
		{
			ID:    token.T_STRING,
			Value: []byte("count"),
		},
		{
			ID:    '(',
			Value: []byte("("),
		},
		{
			ID:    ')',
			Value: []byte(")"),
		},
	}
	suite.Run()
}
