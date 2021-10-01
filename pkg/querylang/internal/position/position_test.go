package position_test

import (
	"testing"

	"github.com/i582/CodeQuery/pkg/querylang/ast"
	builder "github.com/i582/CodeQuery/pkg/querylang/internal/position"
	"github.com/i582/CodeQuery/pkg/querylang/position"
	"github.com/i582/CodeQuery/pkg/querylang/token"
	"gotest.tools/assert"
)

func TestNewTokenPos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewTokenPos(tkn)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 1, EndPos: 3}, pos)
}

func TestNewTokensPos(t *testing.T) {
	token1 := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    6,
		},
	}

	pos := builder.NewBuilder().NewTokensPos(token1, token2)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 2, EndPos: 6}, pos)
}

func TestNewNodePos(t *testing.T) {
	n := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewNodePos(n)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 1, EndPos: 3}, pos)
}

func TestNewTokenNodePos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	n := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    12,
		},
	}

	pos := builder.NewBuilder().NewTokenNodePos(tkn, n)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 2, EndPos: 12}, pos)
}

func TestNewNodeTokenPos(t *testing.T) {
	n := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    12,
		},
	}

	pos := builder.NewBuilder().NewNodeTokenPos(n, tkn)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 2, EndPos: 12}, pos)
}

func TestNewNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	pos := builder.NewBuilder().NewNodeListPos([]ast.Node{n1, n2})

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 2, EndPos: 19}, pos)
}

func TestNewNodesPos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	pos := builder.NewBuilder().NewNodesPos(n1, n2)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 2, EndPos: 19}, pos)
}

func TestNewNodeListTokenPos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 3,
			EndLine:   3,
			StartPos:  20,
			EndPos:    22,
		},
	}

	pos := builder.NewBuilder().NewNodeListTokenPos([]ast.Node{n1, n2}, tkn)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 3, EndPos: 22}, pos)
}

func TestNewTokenNodeListPos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    2,
		},
	}

	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  3,
			EndPos:    10,
		},
	}

	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 3,
			EndLine:   3,
			StartPos:  11,
			EndPos:    20,
		},
	}

	pos := builder.NewBuilder().NewTokenNodeListPos(tkn, []ast.Node{n1, n2})

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 3, EndPos: 20}, pos)
}

func TestNewNodeNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}

	n3 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	pos := builder.NewBuilder().NewNodeNodeListPos(n1, []ast.Node{n2, n3})

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 3, EndPos: 26}, pos)
}

func TestNewNodeListNodePos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}
	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}
	n3 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	pos := builder.NewBuilder().NewNodeListNodePos([]ast.Node{n1, n2}, n3)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 3, EndPos: 26}, pos)
}

func TestNewOptionalListTokensPos(t *testing.T) {
	token1 := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    6,
		},
	}

	pos := builder.NewBuilder().NewOptionalListTokensPos(nil, token1, token2)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: 2, EndPos: 6}, pos)
}

func TestNewOptionalListTokensPos2(t *testing.T) {
	n2 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}
	n3 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	token1 := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 4,
			EndLine:   4,
			StartPos:  27,
			EndPos:    29,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 5,
			EndLine:   5,
			StartPos:  30,
			EndPos:    32,
		},
	}

	pos := builder.NewBuilder().NewOptionalListTokensPos([]ast.Node{n2, n3}, token1, token2)

	assert.DeepEqual(t, &position.Pos{StartLine: 2, EndLine: 5, StartPos: 9, EndPos: 32}, pos)
}

func TestNilNodePos(t *testing.T) {
	pos := builder.NewBuilder().NewNodesPos(nil, nil)

	assert.DeepEqual(t, &position.Pos{StartLine: -1, EndLine: -1, StartPos: -1, EndPos: -1}, pos)
}

func TestNilNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	pos := builder.NewBuilder().NewNodeNodeListPos(n1, nil)

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: -1, EndPos: -1}, pos)
}

func TestNilNodeListTokenPos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewNodeListTokenPos(nil, tkn)

	assert.DeepEqual(t, &position.Pos{StartLine: -1, EndLine: 1, StartPos: -1, EndPos: 3}, pos)
}

func TestEmptyNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	pos := builder.NewBuilder().NewNodeNodeListPos(n1, []ast.Node{})

	assert.DeepEqual(t, &position.Pos{StartLine: 1, EndLine: -1, EndPos: -1}, pos)
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Pos: &position.Pos{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewBuilder().NewNodeListTokenPos([]ast.Node{}, tkn)

	assert.DeepEqual(t, &position.Pos{StartLine: -1, EndLine: 1, StartPos: -1, EndPos: 3}, pos)
}
