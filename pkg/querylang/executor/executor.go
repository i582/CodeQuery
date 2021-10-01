package executor

import (
	"fmt"
	"strconv"

	models2 "github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	"github.com/i582/CodeQuery/pkg/querylang/executor/filter"
	"github.com/i582/CodeQuery/pkg/querylang/executor/runtime"
	"github.com/i582/CodeQuery/pkg/querylang/internal/query"
	"github.com/i582/CodeQuery/pkg/querylang/models"
	"github.com/i582/CodeQuery/pkg/querylang/token"
	"github.com/i582/CodeQuery/pkg/utils"
)

type Executor struct {
	db *models2.Database

	data map[string]interface{}

	value  interface{}
	errors []string

	cache map[ast.Node]interface{}
}

func NewExecutor(db *models2.Database) *Executor {
	return &Executor{db: db, data: map[string]interface{}{}, cache: map[ast.Node]interface{}{}}
}

func (e *Executor) ExecDepSelect(selectExpr *ast.SelectExpr) interface{} {
	fromNode := selectExpr.FromExpr.From

	switch from := fromNode.(type) {
	case *ast.SelectExpr:
		val := e.ExecSelect(from)
		funcs := val.(models2.FuncTable)
		if len(funcs.Data) != 1 {
			e.fatal("'SELECT' subexpression for 'deps' must return exactly one function, use 'LIMIT 1' or refine the subquery")
		}

		var depth int64 = 3
		if selectExpr.WithExpr != nil {
			for _, itemRaw := range selectExpr.WithExpr.WithList.Items {
				switch item := itemRaw.(type) {
				case *ast.ComparisonExpr:
					if item.OpTkn.ID == token.T_EQUAL {
						if vr, ok := item.Left.(*ast.Variable); ok && string(vr.Name) == "depth" {
							depth = runtime.IntCast(e.Exec(item.Right))
						}
					}
				}
			}
		}

		fun := funcs.Data[0]

		graph := &models2.Graph{
			Root:     fun.ID,
			Graph:    map[int64]models2.NodesIDs{},
			RevGraph: map[int64]models2.NodesIDs{},
		}

		e.funcDeps(fun.ID, depth+1, map[int64]bool{}, graph, selectExpr.WhereExpr)

		var allPaths [][]*models2.Func

		for id, called := range graph.Graph {
			if len(called) == 0 {
				paths := graph.FindPaths(e.db, id, graph.Root, 10)
				graph.ReversePaths(paths)
				allPaths = append(allPaths, paths...)
			}
		}

		filteredGraph := &models2.Graph{
			Root:     fun.ID,
			Graph:    map[int64]models2.NodesIDs{},
			RevGraph: map[int64]models2.NodesIDs{},
		}

		for _, path := range allPaths {
			for i := 0; i < len(path)-1; i++ {
				if selectExpr.WhereExpr != nil {
					e.data["path"] = models2.FuncPath(path)
					selectExpr.WhereExpr.Expr.Walk(e)
					if runtime.BoolCast(e.value) {
						filteredGraph.Graph[path[i].ID] = append(filteredGraph.Graph[path[i].ID], path[i+1].ID)
					}
				} else {
					filteredGraph.Graph[path[i].ID] = append(filteredGraph.Graph[path[i].ID], path[i+1].ID)
				}
			}
		}

		for funID, ids := range filteredGraph.Graph {
			idMap := make(map[int64]struct{}, len(ids))
			for _, id := range ids {
				idMap[id] = struct{}{}
			}
			ids = ids[:0]
			for id := range idMap {
				ids = append(ids, id)
			}
			filteredGraph.Graph[funID] = ids
		}

		return filteredGraph
	}

	return nil
}

func (e *Executor) funcDeps(id int64, depth int64, visited map[int64]bool, graph *models2.Graph, whereNode *ast.WhereExpr) bool {
	if depth == 0 {
		return false
	}

	if needAdd, ok := visited[id]; ok {
		return needAdd
	}
	visited[id] = true

	fun := e.db.GetFuncByID(id)

	graph.Graph[fun.ID] = nil

	callee := e.getFuncCallee(fun)

	for _, funcID := range callee {
		added := e.funcDeps(funcID, depth-1, visited, graph, whereNode)
		if added {
			graph.Graph[fun.ID] = append(graph.Graph[fun.ID], funcID)
			graph.RevGraph[funcID] = append(graph.RevGraph[funcID], fun.ID)
		}
	}

	return true
}

func (e *Executor) getFuncCallee(fun *models2.Func) models2.NodesIDs {
	calleeFuncsIDs := make(models2.NodesIDs, 0)

	for _, graph := range e.db.Graphs {
		calleeIDs := graph.Graph[fun.ID]

		calleeFuncsIDs = append(calleeFuncsIDs, calleeIDs...)
	}

	return calleeFuncsIDs
}

func (e *Executor) ExecSelect(selectExpr *ast.SelectExpr) (res interface{}) {
	if cacheResult, ok := e.cache[selectExpr]; ok {
		return cacheResult
	}
	defer func() {
		e.cache[selectExpr] = res
	}()

	needCount := selectExpr.Select.Count != nil

	fromNode := selectExpr.FromExpr.From

	if selectExpr.Select.List != nil &&
		string(selectExpr.Select.List.Items[0].(*ast.Identifier).Value) == "deps" {
		return e.ExecDepSelect(selectExpr)
	}

	opts := filter.Opts{Limit: 20, OrderField: "id"}

	if selectExpr.LimitExpr != nil {
		opts.Limit = runtime.IntCast(e.Exec(selectExpr.LimitExpr))
	}

	if selectExpr.OrderByExpr != nil {
		vr, ok := selectExpr.OrderByExpr.Field.(*ast.Variable)
		if !ok {
			panic("name of the field to be sorted must be a simple identifier")
		}
		opts.OrderField = string(vr.Name)
		opts.RevOrder = selectExpr.OrderByExpr.Desc
	}

	if selectExpr.Select.List != nil {
		for _, col := range selectExpr.Select.List.Items {
			name := string(col.(*ast.Identifier).Value)
			opts.SelectCols = append(opts.SelectCols, name)
		}
	}

	switch from := fromNode.(type) {

	case *ast.MethodCallExpr:
		val := e.Exec(from)
		switch value := val.(type) {
		case models.IGlobals:
			return e.runSelectExprForGlobals(selectExpr, value.(models2.Globals))
		}

	case *ast.Variable:
		switch string(from.Name) {
		case "funcs":
			data := e.runSelectExprForFuncs(selectExpr, e.db.Funcs)
			filtered := filter.FuncsFilter(data.(models2.FuncMap), opts)
			if needCount {
				return filtered.Length()
			}

			return filtered
		case "globals":
			data := e.runSelectExprForGlobals(selectExpr, e.db.Globals)
			filtered := filter.GlobalsFilter(data.(models2.Globals), opts)
			if needCount {
				return filtered.Length()
			}

			return filtered

		case "calls":
			filtered := make(models2.FuncCalls, 0)

			for _, fun := range e.db.Funcs {
				for _, call := range fun.Calls_ {
					e.data["call"] = call
					e.data["func"] = fun
					if selectExpr.WhereExpr != nil {
						selectExpr.WhereExpr.Expr.Walk(e)
						if runtime.BoolCast(e.value) {
							filtered = append(filtered, call)
						}
					} else {
						filtered = append(filtered, call)
					}
				}
			}

			if needCount {
				return len(filtered)
			}

			return filter.FuncCallsFilter(filtered, opts)
		}
	}

	return nil
}

func (e *Executor) runSelectExprForGlobals(selectExpr *ast.SelectExpr, data models2.Globals) interface{} {
	filtered := make(map[int64]models2.Global, len(data.ByID))

	for _, global := range data.ByID {
		if selectExpr.WhereExpr == nil {
			filtered[global.ID] = global
			continue
		}
		e.data["global"] = global
		selectExpr.WhereExpr.Expr.Walk(e)
		if runtime.BoolCast(e.value) {
			filtered[global.ID] = global
		}
	}

	res := models2.Globals{ByID: filtered}
	e.cache[selectExpr] = res

	return res
}

func (e *Executor) runSelectExprForFuncs(selectExpr *ast.SelectExpr, data models2.FuncMap) interface{} {
	filtered := models2.FuncMap{}

	if selectExpr.WhereExpr != nil {
		filtered = make(models2.FuncMap, len(data))

		for name, fun := range data {
			e.data["func"] = fun
			selectExpr.WhereExpr.Expr.Walk(e)
			if runtime.BoolCast(e.value) {
				filtered[name] = fun
			}
		}
	} else {
		filtered = data
	}

	return filtered
}

func (e *Executor) Exec(n ast.Node) interface{} {
	n.Walk(e)
	return e.value
}

func (e *Executor) fatal(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func (e *Executor) EnterNode(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.SelectExpr:
		e.value = e.ExecSelect(n)
		return false

	case *ast.Variable:
		e.value = e.data[string(n.Name)]

	case *ast.MethodCallExpr:
		name := string(n.MethodName.Value)
		val := e.Exec(n.Variable)

		switch value := val.(type) {
		case models.IFunc:
			switch name {
			case "countUse":
				e.checkCountArgs(n, 0)
				e.value = value.CountUse()
			case "globals":
				e.checkCountArgs(n, 0)
				e.value = value.Globals()
			case "calls":
				e.checkCountArgs(n, 0)
				e.value = value.Calls()
			case "fullName":
				e.checkCountArgs(n, 0)
				e.value = value.FullName()
			case "name":
				e.checkCountArgs(n, 0)
				e.value = value.Name()
			case "namespace":
				e.checkCountArgs(n, 0)
				e.value = value.Namespace()
			case "className":
				e.checkCountArgs(n, 0)
				e.value = value.ClassName()
			default:
				e.fatal("Method '" + name + "' is not defined for 'Func' struct")
			}

		case models.IGlobals:
			switch name {
			case "count":
				e.checkCountArgs(n, 0)
				e.value = value.Count()
			case "contains":
				e.checkCountArgs(n, 1)

				globals := e.Exec(n.Args[0]).(models.IGlobals)
				e.value = value.Contains(globals)
			default:
				e.fatal("Method '" + name + "' is not defined for 'Globals' struct")
			}

		case models.IGlobal:
			switch name {
			case "name":
				e.checkCountArgs(n, 0)
				e.value = value.Name()
			case "countUse":
				e.checkCountArgs(n, 0)
				e.value = value.CountUse()
			default:
				e.fatal("Method '" + name + "' is not defined for 'Global' struct")
			}

		case models.IFuncCalls:
			switch name {
			case "count":
				e.checkCountArgs(n, 0)
				e.value = value.Count()

			default:
				e.fatal("Method '" + name + "' is not defined for 'FuncCalls' struct")
			}

		case models.IFuncCall:
			switch name {
			case "func":
				e.checkCountArgs(n, 0)
				e.value = e.data["func"]

			case "args":
				e.checkCountArgs(n, 0)
				e.value = value.Args()

			case "arg":
				e.checkCountArgs(n, 1)
				e.value = value.Arg(runtime.IntCast(e.Exec(n.Args[0])))

			default:
				e.fatal("Method '" + name + "' is not defined for 'FuncCall' struct")
			}

		case models.IFuncArg:
			switch name {
			case "string":
				e.checkCountArgs(n, 0)
				e.value = value.String()
			case "isInt":
				e.checkCountArgs(n, 0)
				e.value = value.IsInt()
			case "isFloat":
				e.checkCountArgs(n, 0)
				e.value = value.IsFloat()
			case "isBool":
				e.checkCountArgs(n, 0)
				e.value = value.IsBool()
			case "isString":
				e.checkCountArgs(n, 0)
				e.value = value.IsString()
			case "isConstant":
				e.checkCountArgs(n, 0)
				e.value = value.IsConstant()
			case "isVariable":
				e.checkCountArgs(n, 0)
				e.value = value.IsVariable()
			case "isExpression":
				e.checkCountArgs(n, 0)
				e.value = value.IsExpression()

			default:
				e.fatal("Method '" + name + "' is not defined for 'FuncCallArg' struct")
			}

		case models.IFuncPath:
			switch name {
			case "begin":
				e.checkCountArgs(n, 0)
				e.value = value.Begin()

			case "end":
				e.checkCountArgs(n, 0)
				e.value = value.End()

			case "length":
				e.checkCountArgs(n, 0)
				e.value = value.Length()

			case "at":
				e.checkCountArgs(n, 1)
				e.value = value.At(runtime.IntCast(e.Exec(n.Args[0])))

			case "contains":
				e.checkCountArgs(n, 1)

				funcs := e.Exec(n.Args[0]).(models2.FuncTable)
				e.value = value.Contains(&funcs)

			default:
				e.fatal("Method '" + name + "' is not defined for 'FuncPath' struct")
			}

		case string:
			switch name {
			case "contains":
				e.checkCountArgs(n, 1)

				e.value = runtime.StringContains(value, e.Exec(n.Args[0]))
			default:
				e.fatal("Method '" + name + "' is not defined for 'string'")
			}
		}

		return false

	case *ast.ComparisonExpr:
		left := e.Exec(n.Left)
		right := e.Exec(n.Right)

		switch n.OpTkn.ID {
		case query.T_EQUAL:
			e.value = runtime.Equal(left, right)
		case query.T_NOT_EQUAL:
			e.value = !runtime.Equal(left, right)
		case query.T_SMALLER:
			e.value = runtime.Smaller(left, right)
		case query.T_GREATER:
			e.value = runtime.Smaller(right, left)
		}

		return false

	case *ast.BinaryExpr:
		switch n.OpTkn.ID {
		case query.T_AND:
			e.value = false
			if runtime.BoolCast(e.Exec(n.Left)) {
				if runtime.BoolCast(e.Exec(n.Right)) {
					e.value = true
				}
			}

		case query.T_OR:
			e.value = false
			if runtime.BoolCast(e.Exec(n.Left)) || runtime.BoolCast(e.Exec(n.Right)) {
				e.value = true
			}
		}

		return false

	case *ast.NotExpr:
		e.value = !runtime.BoolCast(e.Exec(n.Expr))
		return false

	case *ast.BasicLit:
		switch n.Kind.ID {
		case query.T_LNUMBER:
			value, err := strconv.ParseInt(string(n.Value), 0, 64)
			if err == nil {
				e.value = value
			}

		case query.T_DNUMBER:
			value, err := strconv.ParseFloat(string(n.Value), 64)
			if err != nil {
				e.value = value
			}

		case query.T_CONSTANT_STRING:
			val := string(n.Value)
			if val[0] == '"' || val[0] == '\'' {
				e.value = utils.Unquote(val)
			} else {
				e.value = val
			}

		}
	}

	return true
}

func (e *Executor) checkCountArgs(n *ast.MethodCallExpr, expected int) {
	if len(n.Args) < expected {
		e.fatal("Too few arguments for %s() method, expecting %d, saw %d", n.MethodName.Value, expected, len(n.Args))
	}

	if len(n.Args) > expected {
		e.fatal("Too many arguments for %s() method, expecting %d, saw %d", n.MethodName.Value, expected, len(n.Args))
	}
}

func (e *Executor) LeaveNode(n ast.Node) {
}
