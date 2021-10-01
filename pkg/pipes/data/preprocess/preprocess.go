package preprocess

import (
	"strings"

	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/ir/irutil"
	"github.com/VKCOM/noverify/src/meta"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/pipes/collect/symbols"
	"github.com/i582/CodeQuery/pkg/pipes/collect/walkers"
)

func convertPosition(pos meta.ElementPosition) models.Pos {
	return models.Pos{
		Filename: pos.Filename,
		Line:     pos.Line,
		EndLine:  pos.EndLine,
		Col:      pos.Character,
		Length:   pos.Length,
	}
}

func convertGlobal(vr *symbols.GlobalVar, id int64) models.Global {
	return models.Global{
		ID:       id,
		Name_:    vr.Name,
		UseCount: vr.UseCount,
	}
}

func convertGlobals(vars *symbols.GlobalVars, db *models.Database) models.Globals {
	res := make(map[int64]models.Global, len(vars.GlobalVars))
	for _, globalVar := range vars.GlobalVars {
		global := db.Globals.ByName[globalVar.Name]
		res[global.ID] = global
	}
	return models.Globals{ByID: res}
}

func typeOfArgExpression(n ir.Node) models.FuncCallArgType {
	switch n := n.(type) {
	case *ir.Argument:
		return typeOfArgExpression(n.Expr)
	case *ir.ParenExpr:
		return typeOfArgExpression(n.Expr)
	case *ir.Lnumber:
		return models.IntArg
	case *ir.Dnumber:
		return models.FloatArg
	case *ir.ConstFetchExpr:
		if n.Constant.Value == "true" || n.Constant.Value == "false" {
			return models.BoolArg
		}
		return models.ConstantArg
	case *ir.String:
		return models.StringArg
	case *ir.SimpleVar:
		return models.VariableArg
	}

	return models.ExpressionArg
}

func convertFunctionCallArgs(args []symbols.FunctionArg) []models.FuncCallArg {
	res := make([]models.FuncCallArg, 0, len(args))

	for _, arg := range args {
		res = append(res, models.FuncCallArg{
			Type:  typeOfArgExpression(arg.Expr),
			Value: irutil.FmtNode(arg.Expr),
		})
	}

	return res
}

func convertFunctionCalls(calls *symbols.FunctionCalls) []models.FuncCall {
	res := make([]models.FuncCall, 0, len(calls.Calls))

	for _, call := range calls.Calls {
		res = append(res, models.FuncCall{
			Pos:   convertPosition(call.Position),
			Args_: convertFunctionCallArgs(call.Args),
		})
	}

	return res
}

func convertFunction(fun *symbols.Function, id int64, db *models.Database) *models.Func {
	fqn := convertName(fun.Name)

	return &models.Func{
		ID:       id,
		Fqn:      fqn,
		Pos_:     convertPosition(fun.Pos),
		Rem:      fun.NeedRemove,
		Globals_: convertGlobals(fun.Globals, db),
		Calls_:   convertFunctionCalls(fun.Calls),
		UseCount: int64(len(fun.Calls.Calls)),
	}
}

func convertName(name string) models.FQN {
	fqn := models.FQN{}

	if strings.Contains(name, "::") {
		parts := strings.Split(name, "::")
		fqn.Name_ = parts[1]

		namespaceClass := parts[0]
		if namespaceClass == "" {
			fqn.ClassName_ = "undefinedClass"
			return fqn
		}

		lastSlashIndex := strings.LastIndex(namespaceClass, `\`)
		if lastSlashIndex == -1 {
			lastSlashIndex = 0
		}

		fqn.Namespace_ = namespaceClass[:lastSlashIndex]
		fqn.ClassName_ = namespaceClass[lastSlashIndex+1:]
	} else {
		lastSlashIndex := strings.LastIndex(name, `\`)
		if lastSlashIndex == -1 {
			lastSlashIndex = 0
		}
		fqn.Namespace_ = name[:lastSlashIndex]
		fqn.Name_ = name[lastSlashIndex+1:]
	}

	return fqn
}

func Run(data *walkers.GlobalContext) *models.Database {
	db := &models.Database{
		Funcs: make(map[models.FQN]*models.Func, len(data.Functions.Functions)),
		Globals: models.Globals{
			ByID:   make(map[int64]models.Global, len(data.Globals.GlobalVars)),
			ByName: make(map[string]models.Global, len(data.Globals.GlobalVars)),
		},
	}

	globalIdCounter := int64(0)
	for _, globalVar := range data.Globals.GlobalVars {
		g := convertGlobal(globalVar, globalIdCounter)
		db.Globals.ByName[g.Name()] = g
		db.Globals.ByID[g.Id()] = g
		globalIdCounter++
	}

	funcIdCounter := int64(0)
	for _, function := range data.Functions.Functions {
		if strings.Contains(function.Name, "src$") {
			continue
		}

		f := convertFunction(function, funcIdCounter, db)
		db.Funcs[f.Fqn] = f
		funcIdCounter++
	}

	db.Graphs = NodesToGraphs(FunctionsToNodes(data.Functions, db))

	return db
}
