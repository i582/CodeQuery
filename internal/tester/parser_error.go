package tester

import (
	"testing"

	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/errors"
	"github.com/i582/CodeQuery/pkg/querylang/parser"
	"gotest.tools/assert"
)

type ParserErrorTestSuite struct {
	t *testing.T

	Code     string
	Expected []*errors.Error
}

func NewParserErrorTestSuite(t *testing.T) *ParserErrorTestSuite {
	return &ParserErrorTestSuite{
		t: t,
	}
}

func (p *ParserErrorTestSuite) Run() {
	var errs []*errors.Error

	config := conf.Config{
		ErrorHandlerFunc: func(e *errors.Error) {
			errs = append(errs, e)
		},
	}

	parser.Parse([]byte(p.Code), config)
	assert.DeepEqual(p.t, p.Expected, errs)
}
