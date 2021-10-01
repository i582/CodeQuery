package tester

import (
	"testing"

	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/parser"
	"gotest.tools/assert"
)

type ParserTestSuite struct {
	t *testing.T

	Code     string
	Expected ast.Node
}

func NewParserTestSuite(t *testing.T) *ParserTestSuite {
	return &ParserTestSuite{
		t: t,
	}
}

func (p *ParserTestSuite) Run() {
	actual := parser.Parse([]byte(p.Code), conf.Config{})

	assert.DeepEqual(p.t, p.Expected, actual)
}
