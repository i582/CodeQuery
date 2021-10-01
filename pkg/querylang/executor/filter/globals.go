package filter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/i582/CodeQuery/pkg/models"
)

var defaultGlobalTableCols = map[string]models.TableCol{
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
}

func GlobalsFilter(globals models.Globals, opts Opts) models.GlobalTable {
	var funcTableCols []models.TableCol

	if opts.SelectCols == nil {
		funcTableCols = []models.TableCol{
			defaultGlobalTableCols["id"],
			defaultGlobalTableCols["name"],
			defaultGlobalTableCols["uses"],
		}
	} else {
		funcTableCols = append(funcTableCols, defaultGlobalTableCols["id"])
		for _, col := range opts.SelectCols {
			name, ok := defaultGlobalTableCols[col]
			if !ok {
				panic(fmt.Sprintf("unknown field '%s' for the 'Global' structure", col))
			}
			funcTableCols = append(funcTableCols, name)
		}
	}

	globalList := make([]models.Global, 0, len(globals.ByID))
	for _, global := range globals.ByID {
		globalList = append(globalList, global)
	}

	sort.Slice(globalList, func(i, j int) bool {
		var fun1 int64
		var fun2 int64
		switch opts.OrderField {
		case "id":
			fun1 = globalList[i].ID
			fun2 = globalList[j].ID
		case "name":
			fun1Name := strings.ToLower(globalList[i].Name())
			fun2Name := strings.ToLower(globalList[j].Name())
			if opts.RevOrder {
				fun1Name, fun2Name = fun2Name, fun1Name
			}
			return fun1Name < fun2Name
		case "uses":
			fun1 = globalList[i].CountUse()
			fun2 = globalList[j].CountUse()

		default:
			panic(fmt.Sprintf("unknown field '%s' for the 'Func' structure", opts.OrderField))
		}

		if opts.RevOrder {
			fun1, fun2 = fun2, fun1
		}

		return fun1 < fun2
	})

	if opts.Limit != -1 && int64(len(globalList)) > opts.Limit {
		return models.GlobalTable{Data: globalList[:opts.Limit], Cols: funcTableCols}
	}

	return models.GlobalTable{Data: globalList, Cols: funcTableCols}
}
