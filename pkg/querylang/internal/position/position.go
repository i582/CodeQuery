package position

import (
	"github.com/i582/CodeQuery/pkg/querylang/ast"
	position2 "github.com/i582/CodeQuery/pkg/querylang/position"
	"github.com/i582/CodeQuery/pkg/querylang/token"
)

type startPos struct {
	startLine int
	startPos  int
}

type endPos struct {
	endLine int
	endPos  int
}

type Builder struct {
	pool *position2.Pool
}

func NewBuilder() *Builder {
	return &Builder{
		pool: position2.NewPool(position2.DefaultBlockSize),
	}
}

func getListStartPos(l []ast.Node) startPos {
	if l == nil {
		return startPos{-1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1}
	}

	return getNodeStartPos(l[0])
}

func getNodeStartPos(n ast.Node) startPos {
	sl := -1
	sp := -1

	if n == nil {
		return startPos{-1, -1}
	}

	// p := n.GetPos()
	// if p != nil {
	// 	sl = p.StartLine
	// 	sp = p.StartPos
	// }

	return startPos{sl, sp}
}

func getListEndPos(l []ast.Node) endPos {
	if l == nil {
		return endPos{-1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1}
	}

	return getNodeEndPos(l[len(l)-1])
}

func getNodeEndPos(n ast.Node) endPos {
	el := -1
	ep := -1

	if n == nil {
		return endPos{-1, -1}
	}

	// p := n.GetPos()
	// if p != nil {
	// 	el = p.EndLine
	// 	ep = p.EndPos
	// }

	return endPos{el, ep}
}

// NewNodeListPos returns new Pos_
func (b *Builder) NewNodeListPos(list []ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = getListEndPos(list).endPos

	return pos
}

// NewNodePos returns new Pos_
func (b *Builder) NewNodePos(n ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(n).startLine
	pos.EndLine = getNodeEndPos(n).endLine
	pos.StartPos = getNodeStartPos(n).startPos
	pos.EndPos = getNodeEndPos(n).endPos

	return pos
}

// NewTokenPos returns new Pos_
func (b *Builder) NewTokenPos(t *token.Token) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = t.Pos.StartLine
	pos.EndLine = t.Pos.EndLine
	pos.StartPos = t.Pos.StartPos
	pos.EndPos = t.Pos.EndPos

	return pos
}

// NewTokensPos returns new Pos_
func (b *Builder) NewTokensPos(startToken *token.Token, endToken *token.Token) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = startToken.Pos.StartLine
	pos.EndLine = endToken.Pos.EndLine
	pos.StartPos = startToken.Pos.StartPos
	pos.EndPos = endToken.Pos.EndPos

	return pos
}

// NewTokenNodePos returns new Pos_
func (b *Builder) NewTokenNodePos(t *token.Token, n ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = t.Pos.StartLine
	pos.EndLine = getNodeEndPos(n).endLine
	pos.StartPos = t.Pos.StartPos
	pos.EndPos = getNodeEndPos(n).endPos

	return pos
}

// NewNodeTokenPos returns new Pos_
func (b *Builder) NewNodeTokenPos(n ast.Node, t *token.Token) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(n).startLine
	pos.EndLine = t.Pos.EndLine
	pos.StartPos = getNodeStartPos(n).startPos
	pos.EndPos = t.Pos.EndPos

	return pos
}

// NewNodesPos returns new Pos_
func (b *Builder) NewNodesPos(startNode ast.Node, endNode ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(startNode).startLine
	pos.EndLine = getNodeEndPos(endNode).endLine
	pos.StartPos = getNodeStartPos(startNode).startPos
	pos.EndPos = getNodeEndPos(endNode).endPos

	return pos
}

// NewNodeListTokenPos returns new Pos_
func (b *Builder) NewNodeListTokenPos(list []ast.Node, t *token.Token) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = t.Pos.EndLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = t.Pos.EndPos

	return pos
}

// NewTokenNodeListPos returns new Pos_
func (b *Builder) NewTokenNodeListPos(t *token.Token, list []ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = t.Pos.StartLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = t.Pos.StartPos
	pos.EndPos = getListEndPos(list).endPos

	return pos
}

// NewNodeNodeListPos returns new Pos_
func (b *Builder) NewNodeNodeListPos(n ast.Node, list []ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getNodeStartPos(n).startLine
	pos.EndLine = getListEndPos(list).endLine
	pos.StartPos = getNodeStartPos(n).startPos
	pos.EndPos = getListEndPos(list).endPos

	return pos
}

// NewNodeListNodePos returns new Pos_
func (b *Builder) NewNodeListNodePos(list []ast.Node, n ast.Node) *position2.Pos {
	pos := b.pool.Get()

	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = getNodeEndPos(n).endLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = getNodeEndPos(n).endPos

	return pos
}

// NewOptionalListTokensPos returns new Pos_
func (b *Builder) NewOptionalListTokensPos(list []ast.Node, t *token.Token, endToken *token.Token) *position2.Pos {
	pos := b.pool.Get()

	if list == nil {
		pos.StartLine = t.Pos.StartLine
		pos.EndLine = endToken.Pos.EndLine
		pos.StartPos = t.Pos.StartPos
		pos.EndPos = endToken.Pos.EndPos

		return pos
	}
	pos.StartLine = getListStartPos(list).startLine
	pos.EndLine = endToken.Pos.EndLine
	pos.StartPos = getListStartPos(list).startPos
	pos.EndPos = endToken.Pos.EndPos

	return pos
}
