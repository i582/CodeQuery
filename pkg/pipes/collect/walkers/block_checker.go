package walkers

import (
	"sync/atomic"

	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/linter"
	"github.com/i582/CodeQuery/pkg/pipes/collect/symbols"
)

// BlockIndexer is a dummy walker.
type BlockIndexer struct {
	linter.BlockCheckerDefaults
}

// BlockChecker is a walker that handles function calls, method calls,
// class creation, and file imports in block context (inside functions).
type BlockChecker struct {
	linter.BlockCheckerDefaults
	ctx  *linter.BlockContext
	root *RootChecker
}

// NewBlockChecker creates a new BlockChecker walker.
func NewBlockChecker(ctx *linter.BlockContext, root *RootChecker) *BlockChecker {
	return &BlockChecker{
		ctx:  ctx,
		root: root,
	}
}

// EnterNode is method to use BlockChecker in the Walk method of AST nodes.
func (b *BlockChecker) EnterNode(n ir.Node) bool {
	b.AfterEnterNode(n)
	return true
}

// LeaveNode is method to use BlockChecker in the Walk method of AST nodes.
func (b *BlockChecker) LeaveNode(n ir.Node) {}

// AfterEnterNode is the main method for processing AST nodes.
func (b *BlockChecker) AfterEnterNode(n ir.Node) {
	switch n := n.(type) {
	case *ir.GlobalStmt:
		curFunc, ok := b.root.getCurrentFunc()
		if !ok {
			return
		}

		for _, variable := range n.Vars {
			vr, ok := variable.(*ir.SimpleVar)
			if !ok {
				continue
			}

			globalVar, ok := b.root.globalCtx.Globals.Get(vr.Name)
			if !ok {
				globalVar = &symbols.GlobalVar{
					Name:     vr.Name,
					UseCount: 1,
				}
				curFunc.Globals.Add(globalVar)
				b.root.globalCtx.Globals.Add(globalVar)
			} else {
				atomic.AddInt64(&globalVar.UseCount, 1)
			}
			curFunc.Globals.Add(globalVar)
		}

	case *ir.NewExpr:
		b.root.handleNew(n, b.ctx.Scope())
	case *ir.FunctionCallExpr:
		b.root.handleFunctionCall(n, b.ctx.Scope(), b)
	case *ir.StaticCallExpr:
		b.root.handleStaticCall(n, b.ctx.Scope())
	case *ir.MethodCallExpr:
		b.root.handleMethodCall(n, b.ctx.Scope(), b)
	case *ir.NullsafeMethodCallExpr:
		b.root.handleNullsafeMethodCall(n, b.ctx.Scope(), b)
	case *ir.ImportExpr:
		b.root.handleImportExpr(n)
	case *ir.CloneExpr:
		b.root.handleCloneExpr(n, b.ctx.Scope())
	case *ir.PropertyFetchExpr:
		b.root.handlePropertyFetch(n, b.ctx.Scope(), b.ctx.NodePath())

	case *ir.Assign:
		// Because of the way we handle assignments,
		// we have to redirect our walker explicitly.
		n.Variable.Walk(b)
		n.Expr.Walk(b)
	}
}
