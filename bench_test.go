package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"testing"

	"github.com/i582/CodeQuery/cmd"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/executor"
	"github.com/i582/CodeQuery/pkg/querylang/parser"
)

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd.Run()
	}
}

func BenchmarkRegex(b *testing.B) {
	f, _ := os.OpenFile("1.graph", os.O_RDONLY, 0777)

	db := &models.Database{}
	dec := gob.NewDecoder(f)
	err := dec.Decode(db)
	if err != nil {
		fmt.Println(err)
	}

	f.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root := parser.Parse([]byte(`
/**
 * Comment
 */

-- Comment

SELECT * FROM @funcs WHERE

func.name().contains("debug111") AND 
func.globals().contains(
	SELECT * FROM @globals WHERE global.countUse() = 1
)

`), conf.Config{})

		stmt := root.(*ast.Root).Stmts[0]

		executor.NewExecutor(db).ExecSelect(stmt.(*ast.SelectExpr))
		// for _, fun := range res.(models.FuncMap) {
		// 	fmt.Println(fun.Name_(), fun.CountUse(), fun.ByID().Count(), fun.ByID())
		// }
	}
}
