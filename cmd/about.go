package cmd

import (
	"fmt"

	"github.com/i582/CodeQuery/pkg/shell"
	"github.com/i582/CodeQuery/pkg/shell/flags"
)

func About() *shell.Executor {
	aboutExecutor := &shell.Executor{
		Name:  "about",
		Help:  "shows brief information about CodeQuery",
		Flags: flags.NewFlags(),
		Func: func(c *shell.Context) {
			fmt.Print(`About CodeQuery v0.0.1

CodeQuery is a tool for searching and aggregating data for PHP.

Petr Makhnev MIT (c) 2021
`)
		},
	}

	return aboutExecutor
}
