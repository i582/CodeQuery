package parser

import (
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/internal/query"
)

func Parse(src []byte, config conf.Config) ast.Node {
	lexer := query.NewLexer(src, config)
	parser := query.NewParser(lexer)
	parser.Parse()
	return parser.GetRootNode()
}
