package filter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/i582/CodeQuery/pkg/models"
)

var defaultFuncCallTableCols = map[string]models.TableCol{
	"args": {
		Name:   "Args",
		IdName: "args",
	},
	"file": {
		Name:   "File",
		IdName: "file",
	},
	"args_count": {
		Name:   "Args count",
		IdName: "args_count",
	},
}

func FuncCallsFilter(funcs models.FuncCalls, opts Opts) models.FuncCallTable {
	var funcTableCols []models.TableCol

	if opts.SelectCols == nil {
		funcTableCols = []models.TableCol{
			defaultFuncCallTableCols["args"],
			defaultFuncCallTableCols["args_count"],
			defaultFuncCallTableCols["file"],
		}
	} else {
		for _, col := range opts.SelectCols {
			name, ok := defaultFuncTableCols[col]
			if !ok {
				panic(fmt.Sprintf("unknown field '%s' for the 'FuncCall' structure", col))
			}
			funcTableCols = append(funcTableCols, name)
		}
	}

	sort.Slice(funcs, func(i, j int) bool {
		var fun1 int64
		var fun2 int64
		switch opts.OrderField {
		case "args_count":
			fun1 = funcs[i].Args_()
			fun2 = funcs[j].Args_()
		case "file":
			fun1Name := strings.ToLower(funcs[i].File())
			fun2Name := strings.ToLower(funcs[j].File())
			if opts.RevOrder {
				fun1Name, fun2Name = fun2Name, fun1Name
			}
			return fun1Name < fun2Name
		case "args":
			return fun1 < fun2
		case "id":

		default:
			panic(fmt.Sprintf("unknown field '%s' for the 'FuncCall' structure", opts.OrderField))
		}

		if opts.RevOrder {
			fun1, fun2 = fun2, fun1
		}

		return fun1 < fun2
	})

	if opts.Limit != -1 && int64(len(funcs)) > opts.Limit {
		funcs = funcs[:opts.Limit]
	}

	return models.FuncCallTable{Data: funcs, Cols: funcTableCols}
}
