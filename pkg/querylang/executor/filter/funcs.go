package filter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/i582/CodeQuery/pkg/models"
)

type Opts struct {
	Limit      int64
	OrderField string
	RevOrder   bool
	SelectCols []string
}

var defaultFuncTableCols = map[string]models.TableCol{
	"id": {
		Name:   "#",
		IdName: "id",
	},
	"name": {
		Name:   "Name",
		IdName: "name",
	},
	"uses": {
		Name:   "Count Use",
		IdName: "uses",
	},
	"globals": {
		Name:   "Globals",
		IdName: "globals",
	},
}

func FuncsFilter(funcs models.FuncMap, opts Opts) models.FuncTable {
	var funcTableCols []models.TableCol

	if opts.SelectCols == nil {
		funcTableCols = []models.TableCol{
			defaultFuncTableCols["id"],
			defaultFuncTableCols["name"],
			defaultFuncTableCols["uses"],
			defaultFuncTableCols["globals"],
		}
	} else {
		funcTableCols = append(funcTableCols, defaultFuncTableCols["id"])
		for _, col := range opts.SelectCols {
			name, ok := defaultFuncTableCols[col]
			if !ok {
				panic(fmt.Sprintf("unknown field '%s' for the 'Func' structure", col))
			}
			funcTableCols = append(funcTableCols, name)
		}
	}

	funcList := make([]*models.Func, 0, len(funcs))
	for _, fun := range funcs {
		funcList = append(funcList, fun)
	}

	sort.Slice(funcList, func(i, j int) bool {
		var fun1 int64
		var fun2 int64
		switch opts.OrderField {
		case "id":
			fun1 = funcList[i].ID
			fun2 = funcList[j].ID
		case "name":
			fun1Name := strings.ToLower(funcList[i].Fqn.String())
			fun2Name := strings.ToLower(funcList[j].Fqn.String())
			if opts.RevOrder {
				fun1Name, fun2Name = fun2Name, fun1Name
			}
			return fun1Name < fun2Name

		case "globals":
			fun1 = funcList[i].Globals().Count()
			fun2 = funcList[j].Globals().Count()
		case "uses":
			fun1 = funcList[i].Calls().Count()
			fun2 = funcList[j].Calls().Count()

		default:
			panic(fmt.Sprintf("unknown field '%s' for the 'Func' structure", opts.OrderField))
		}

		if opts.RevOrder {
			fun1, fun2 = fun2, fun1
		}

		return fun1 < fun2
	})

	if opts.Limit != -1 && int64(len(funcList)) > opts.Limit {
		return models.FuncTable{Data: funcList[:opts.Limit], Cols: funcTableCols}
	}

	return models.FuncTable{Data: funcList, Cols: funcTableCols}
}
