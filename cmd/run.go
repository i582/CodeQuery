package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/i582/CodeQuery/pkg/grapher"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/errors"
	"github.com/i582/CodeQuery/pkg/querylang/executor"
	"github.com/i582/CodeQuery/pkg/querylang/parser"
	"github.com/i582/CodeQuery/pkg/representator"
	"github.com/i582/CodeQuery/pkg/shell"
	"github.com/i582/CodeQuery/pkg/shell/flags"
	"github.com/i582/CodeQuery/pkg/utils"
	"github.com/i582/cfmt"
)

func RunQuery() *shell.Executor {
	log.SetFlags(0)

	runExecutor := &shell.Executor{
		Name:      "run",
		Help:      "run query",
		CountArgs: -1,
		Flags: flags.NewFlags(
			&flags.Flag{
				Name:      "f",
				WithValue: true,
				Help:      "file to run",
				Default:   "",
			},
			&flags.Flag{
				Name:      "graph",
				WithValue: true,
				Help:      "name of the generated graph",
				Default:   "",
			},
		),
		Func: func(c *shell.Context) {
			if len(c.Args) == 0 && c.GetFlagValue("f") == "" {
				c.Error("'run' command must have 1 argument, the query to be executed")
				return
			}

			var queryString string

			filePath := c.GetFlagValue("f")
			if filePath != "" {
				c.Args = append(c.Args, "<from file>")
				content, err := os.ReadFile(filePath)
				if err != nil {
					c.Error(fmt.Sprintf("can't opening file with query named %s.db: %v", filePath, err))
					return
				}

				queryString = string(content)
			} else {
				queryString = utils.Unquote(c.Args[0])
			}

			var parseErrors []*errors.Error
			cfg := conf.Config{ErrorHandlerFunc: func(e *errors.Error) {
				parseErrors = append(parseErrors, e)
			}}
			root := parser.Parse([]byte(queryString), cfg)

			if len(parseErrors) > 0 {
				c.Error("query parsing error")
				for _, err := range parseErrors {
					c.Error(err.String())
				}
				return
			}

			if root == nil || len(root.(*ast.Root).Stmts) == 0 {
				c.Error("empty query")
				return
			}

			stmt := root.(*ast.Root).Stmts[0]

			defer func() {
				if err := recover(); err != nil {
					cfmt.Printf("{{Error while executing query:}}::bold %s\n", err)
				}
			}()

			res := executor.NewExecutor(AppCtx.Database).ExecSelect(stmt.(*ast.SelectExpr))
			switch result := res.(type) {
			case int64:
				fmt.Println(result)
			case models.FuncTable:
				fmt.Println(representator.GetTableFunctionsRepr(&result))

			case models.GlobalTable:
				fmt.Println(representator.GetTableGlobalsRepr(&result))

			case models.FuncCallTable:
				fmt.Println(representator.GetTableFunctionsCallsRepr(&result))

			case *models.Graph:
				graphFilePath := c.GetFlagValue("graph")
				if graphFilePath == "" {
					representator.PrintGraph(AppCtx.Database, result)
					return
				}

				g := grapher.NewGrapher(AppCtx.Database)
				graphData := g.FunctionDeps(result)

				err := grapher.WriteGraph(graphFilePath, graphData)
				if err != nil {
					c.Error(err.Error())
					return
				}
				err = grapher.AddGraphScalability(graphFilePath)
				if err != nil {
					c.Error(err.Error())
					return
				}

				cfmt.Printf("{{Graph was successfully saved in SVG format to a file named '%s.svg'}}::green\n", graphFilePath)
			}
		},
	}

	return runExecutor
}
