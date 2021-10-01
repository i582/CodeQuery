package tester

import (
	"testing"
)

type ParserPrintTestSuite struct {
	t *testing.T
}

//
// func NewParserPrintTestSuite(t *testing.T) *ParserPrintTestSuite {
// 	return &ParserPrintTestSuite{
// 		t: t,
// 	}
// }
//
// func (p *ParserPrintTestSuite) Run(code string) {
// 	actual := p.print(p.parse(code))
// 	assert.DeepEqual(p.t, code, actual)
// }
//
// func (p *ParserPrintTestSuite) parse(src string) ast.Node {
// 	root := parser.Parse([]byte(src), conf.Config{})
// 	return root
// }
//
// func (p *ParserPrintTestSuite) print(n ast.Node) string {
// 	o := bytes.NewBufferString("")
//
// 	pr := printer.NewPrinter(o)
// 	n.Accept(pr)
//
// 	return o.String()
// }
