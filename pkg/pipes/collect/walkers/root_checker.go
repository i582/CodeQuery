package walkers

import (
	"path/filepath"
	"strings"

	"github.com/VKCOM/noverify/src/constfold"
	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/ir/irutil"
	"github.com/VKCOM/noverify/src/linter"
	"github.com/VKCOM/noverify/src/meta"
	"github.com/VKCOM/noverify/src/solver"
	"github.com/VKCOM/noverify/src/types"
	"github.com/i582/CodeQuery/pkg/pipes/collect/symbols"
	"github.com/i582/CodeQuery/pkg/pipes/collect/walkers/namegen"
)

// RootChecker is a walker that collects information about the
// colors of functions and checks them for correctness.
type RootChecker struct {
	linter.RootCheckerDefaults

	ctx       *linter.RootContext
	state     *meta.ClassParseState
	globalCtx *GlobalContext

	fileFunction *symbols.Function

	colorTag string
}

// NewRootChecker returns a new walker.
func NewRootChecker(globalCtx *GlobalContext, ctx *linter.RootContext, colorTag string) *RootChecker {
	return &RootChecker{
		ctx:       ctx,
		globalCtx: globalCtx,
		colorTag:  colorTag,
		state:     ctx.ClassParseState(),
	}
}

// EnterNode is method to use RootChecker in the Walk method of AST nodes.
func (r *RootChecker) EnterNode(n ir.Node) bool {
	r.BeforeEnterNode(n)
	return true
}

// LeaveNode is method to use RootChecker in the Walk method of AST nodes.
func (r *RootChecker) LeaveNode(ir.Node) {}

// BeforeEnterFile sets the current function of the file.
func (r *RootChecker) BeforeEnterFile() {
	fileFunctionName := namegen.FileFunction(r.ctx.Filename())
	fun, ok := r.globalCtx.Functions.Get(fileFunctionName)
	if !ok {
		r.fileFunction = &symbols.Function{
			Name:           namegen.FileFunction(r.ctx.Filename()),
			Type:           symbols.MainFunc,
			Globals:        symbols.NewGlobals(),
			Calls:          symbols.NewFunctionCalls(),
			FunctionsCalls: symbols.NewFunctionCalls(),
			Called:         symbols.NewFunctions(),
			CalledBy:       symbols.NewFunctions(),
		}
		return
	}

	r.fileFunction = fun
}

// AfterEnterNode
func (r *RootChecker) AfterEnterNode(n ir.Node) {
	switch n := n.(type) {
	case *ir.NewExpr:
		r.handleNew(n, nil)
	case *ir.CloneExpr:
		r.handleCloneExpr(n, nil)
	case *ir.FunctionCallExpr:
		r.handleFunctionCall(n, nil, r)
	case *ir.StaticCallExpr:
		r.handleStaticCall(n, nil)
	case *ir.MethodCallExpr:
		r.handleMethodCall(n, nil, r)
	case *ir.NullsafeMethodCallExpr:
		r.handleNullsafeMethodCall(n, nil, r)
	case *ir.PropertyFetchExpr:
		r.handlePropertyFetch(n, nil, irutil.NodePath{})
	case *ir.ImportExpr:
		r.handleImportExpr(n)
	}
}

func (r *RootChecker) handlePropertyFetch(n *ir.PropertyFetchExpr, blockScope *meta.Scope, nodePath irutil.NodePath) {
	var propInfo solver.FindPropertyResult
	var propertyName string
	var ok bool

	switch prop := n.Property.(type) {
	case *ir.Identifier:
		propertyName = prop.Value
	default:
		return
	}

	scope := blockScope
	if scope == nil {
		scope = r.ctx.Scope()
	}

	classTypes := solver.ExprType(scope, r.state, n.Variable)
	classesWithoutProp := make([]string, 0, classTypes.Len())

	classTypes.Iterate(func(classType string) {
		if !types.IsClass(classType) {
			return
		}

		propInfo, ok = solver.FindProperty(r.state.Info, classType, propertyName)
		if !ok || (ok && propInfo.Info.IsFromAnnotation()) {
			classesWithoutProp = append(classesWithoutProp, classType)
		}
	})

	if len(classesWithoutProp) == 0 {
		return
	}

	methodName := "__get"
	if r.inAssign(nodePath) {
		methodName = "__set"
	}

	for _, className := range classesWithoutProp {
		fqn := namegen.Method(className, methodName)

		calledFunc, ok := r.globalCtx.Functions.Get(fqn)
		if !ok {
			return
		}

		r.createEdgeWithCurrent(calledFunc)
	}
}

func (r *RootChecker) handleCloneExpr(n *ir.CloneExpr, blockScope *meta.Scope) {
	scope := blockScope
	if scope == nil {
		scope = r.ctx.Scope()
	}

	classTypes := solver.ExprType(scope, r.state, n.Expr)
	classTypes.Iterate(func(classType string) {
		if !types.IsClass(classType) {
			return
		}

		methodInfo, ok := solver.FindMethod(r.state.Info, classType, "__clone")
		if !ok {
			return
		}

		methodName := namegen.Method(methodInfo.ImplName(), "__clone")

		calledFunc, ok := r.globalCtx.Functions.Get(methodName)
		if !ok {
			return
		}

		r.createEdgeWithCurrent(calledFunc)
	})
}

func (r *RootChecker) handleImportExpr(n *ir.ImportExpr) {
	pathValue := constfold.Eval(r.state, n.Expr)
	if !pathValue.IsValid() {
		return
	}

	path, ok := pathValue.ToString()
	if !ok {
		return
	}

	path, ok = r.getImportAbsPath(path)
	if !ok {
		return
	}

	fileFunc, ok := r.globalCtx.Functions.Get(namegen.FileFunction(path))
	if !ok {
		return
	}

	r.createEdgeWithCurrent(fileFunc)
}

func (r *RootChecker) handleFunctionCall(n *ir.FunctionCallExpr, blockScope *meta.Scope, v ir.Visitor) {
	for _, arg := range n.Args {
		arg.Walk(v)
	}

	fqName, ok := solver.GetFuncName(r.state, n.Function)
	if !ok {
		r.asInvokeMethod(n, blockScope)
		return
	}

	calledFunc, ok := r.globalCtx.Functions.Get(fqName)
	if !ok {
		return
	}

	r.createEdgeWithCurrentWithArgs(n, n.Args, calledFunc)
}

func (r *RootChecker) asInvokeMethod(n *ir.FunctionCallExpr, blockScope *meta.Scope) {
	scope := blockScope
	if scope == nil {
		scope = r.ctx.Scope()
	}

	classTypes := solver.ExprType(scope, r.state, n.Function)
	classTypes.Iterate(func(classType string) {
		if !types.IsClass(classType) {
			return
		}

		methodInfo, ok := solver.FindMethod(r.state.Info, classType, "__invoke")
		if !ok {
			return
		}

		methodName := namegen.Method(methodInfo.ImplName(), "__invoke")

		calledFunc, ok := r.globalCtx.Functions.Get(methodName)
		if !ok {
			return
		}

		r.createEdgeWithCurrent(calledFunc)
	})
}

func (r *RootChecker) handleStaticCall(n *ir.StaticCallExpr, blockScope *meta.Scope) {
	method, ok := n.Call.(*ir.Identifier)
	if !ok {
		return
	}
	methodName := method.Value

	scope := blockScope
	if scope == nil {
		scope = r.ctx.Scope()
	}

	var classType types.Map

	if vr, ok := n.Class.(*ir.SimpleVar); ok {
		classType = solver.ExprType(scope, r.state, vr)
	} else {
		className, ok := solver.GetClassName(r.state, n.Class)
		if !ok {
			return
		}

		classType = types.NewMap(className)
	}

	r.handleMethod(n, n.Args, methodName, classType, true)
}

func (r *RootChecker) handleMethodCall(n *ir.MethodCallExpr, blockScope *meta.Scope, v ir.Visitor) {
	method, ok := n.Method.(*ir.Identifier)
	if !ok {
		return
	}
	methodName := method.Value

	scope := blockScope
	if scope == nil {
		scope = r.ctx.Scope()
	}

	classType := solver.ExprType(scope, r.state, n.Variable)

	r.handleMethod(n, n.Args, methodName, classType, false)

	for _, nn := range n.Args {
		nn.Walk(v)
	}
}

func (r *RootChecker) handleNullsafeMethodCall(n *ir.NullsafeMethodCallExpr, blockScope *meta.Scope, v ir.Visitor) {
	method, ok := n.Method.(*ir.Identifier)
	if !ok {
		return
	}
	methodName := method.Value

	scope := blockScope
	if scope == nil {
		scope = r.ctx.Scope()
	}

	classType := solver.ExprType(scope, r.state, n.Variable)

	r.handleMethod(n, n.Args, methodName, classType, false)

	for _, nn := range n.Args {
		nn.Walk(v)
	}
}

func (r *RootChecker) handleNew(n *ir.NewExpr, blockScope *meta.Scope) {
	className, ok := solver.GetClassName(r.state, n.Class)
	if !ok {
		// If we cannot get the name, then we will try to find the
		// type of the expression, and if it is a class, then we
		// will assume that the constructor of this class is called.

		scope := blockScope
		if scope == nil {
			scope = r.ctx.Scope()
		}

		typ := solver.ExprType(scope, r.state, n.Class)
		typ.Iterate(func(classType string) {
			if !types.IsClass(classType) {
				return
			}

			r.handleMethod(n, n.Args, "__construct", types.NewMap(classType), false)
		})

		return
	}

	classType := types.NewMap(className)

	r.handleMethod(n, n.Args, "__construct", classType, false)
}

func (r *RootChecker) handleMethod(n ir.Node, args []ir.Node, methodName string, classTypes types.Map, static bool) {
	classesWithoutMethod := make([]string, 0, classTypes.Len())

	classTypes.Iterate(func(classType string) {
		if !types.IsClass(classType) {
			return
		}

		methodInfo, ok := solver.FindMethod(r.state.Info, classType, methodName)
		if !ok || (ok && methodInfo.Info.IsFromAnnotation()) {
			// If the method is described in the annotation for the class,
			// then it will be found, but in fact it does not exist and must
			// be redirected to the call to __call or __callStatic.
			classesWithoutMethod = append(classesWithoutMethod, classType)
		}

		methodFQN := namegen.Method(methodInfo.ImplName(), methodName)
		if !ok && methodName == "__construct" {
			methodFQN = namegen.DefaultConstructor(classType)
		}

		calledFunc, ok := r.globalCtx.Functions.Get(methodFQN)
		if !ok {
			return
		}

		r.createEdgeWithCurrentWithArgs(n, args, calledFunc)
	})

	r.handleClassWithoutMethod(static, classesWithoutMethod)
}

func (r *RootChecker) handleClassWithoutMethod(static bool, classesWithoutMethod []string) {
	magicMethodName := "__call"
	if static {
		magicMethodName = "__callStatic"
	}

	for _, className := range classesWithoutMethod {
		fqn := namegen.Method(className, magicMethodName)

		calledFunc, ok := r.globalCtx.Functions.Get(fqn)
		if !ok {
			continue
		}

		r.createEdgeWithCurrent(calledFunc)
	}
}

func (r *RootChecker) getImportAbsPath(path string) (string, bool) {
	if filepath.IsAbs(path) {
		return filepath.Clean(path), true
	}

	// If relative path.
	if strings.HasPrefix(path, ".") || strings.HasPrefix(path, "..") {
		currentFilePath := r.ctx.Filename()
		dir := filepath.Dir(currentFilePath)

		absPath := filepath.Clean(filepath.Join(dir, path))
		return absPath, true
	}

	return "", false
}

func (r *RootChecker) getCurrentFunc() (*symbols.Function, bool) {
	name := r.state.CurrentFunction
	if name == "" {
		return r.fileFunction, true
	}

	if r.state.CurrentClass != "" {
		className, ok := solver.GetClassName(r.state, &ir.Name{Value: r.state.CurrentClass})
		if !ok {
			return nil, false
		}

		fn, ok := r.globalCtx.Functions.Get(className + "::" + name)
		if !ok {
			return nil, false
		}

		return fn, true
	}

	funcName, ok := solver.GetFuncName(r.state, &ir.Name{
		Value: name,
	})
	if !ok {
		return nil, false
	}

	fn, ok := r.globalCtx.Functions.Get(funcName)
	if !ok {
		return nil, false
	}

	return fn, true
}

var WithCalls = true

func (r *RootChecker) createEdgeWithCurrentWithArgs(fun ir.Node, args []ir.Node, calledFunc *symbols.Function) {
	curFunc, ok := r.getCurrentFunc()
	if !ok {
		return
	}

	callArgs := make([]symbols.FunctionArg, 0, len(args))
	for _, arg := range args {
		callArgs = append(callArgs, symbols.FunctionArg{
			Expr: arg,
		})
	}

	pos := r.getElementPos(fun)
	call := symbols.FunctionCall{
		Function: calledFunc,
		Position: pos,
		Args:     callArgs,
	}

	if WithCalls {
		calledFunc.Calls.Add(call)
	}

	curFunc.FunctionsCalls.Add(call)
	curFunc.Called.Add(calledFunc)
	calledFunc.CalledBy.Add(curFunc)
}

func (r *RootChecker) createEdgeWithCurrent(calledFunc *symbols.Function) {
	curFunc, ok := r.getCurrentFunc()
	if !ok {
		return
	}

	curFunc.Called.Add(calledFunc)
	calledFunc.CalledBy.Add(curFunc)
}

func (r *RootChecker) inAssign(nodePath irutil.NodePath) bool {
	for i := 0; nodePath.NthParent(i) != nil; i++ {
		if irutil.IsAssign(nodePath.NthParent(i)) {
			return true
		}
	}
	return false
}

func (r *RootChecker) getElementPos(n ir.Node) meta.ElementPosition {
	pos := ir.GetPosition(n)

	return meta.ElementPosition{
		Filename:  r.ctx.ClassParseState().CurrentFile,
		Character: int32(0),
		Line:      int32(pos.StartLine),
		EndLine:   int32(pos.EndLine),
		Length:    int32(pos.EndPos - pos.StartPos),
	}
}
