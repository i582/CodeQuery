package query

import (
	"strings"

	"github.com/i582/CodeQuery/pkg/querylang/conf"
	"github.com/i582/CodeQuery/pkg/querylang/errors"
	position2 "github.com/i582/CodeQuery/pkg/querylang/position"
	token2 "github.com/i582/CodeQuery/pkg/querylang/token"
)

type Lexer struct {
	data           []byte
	errHandlerFunc func(*errors.Error)

	p, pe, cs   int
	ts, te, act int
	stack       []int
	top         int

	heredocLabel []byte
	tokenPool    *token2.Pool
	positionPool *position2.Pool
	newLines     NewLines
}

func NewLexer(data []byte, config conf.Config) *Lexer {
	lex := &Lexer{
		data:           data,
		errHandlerFunc: config.ErrorHandlerFunc,

		pe:    len(data),
		stack: make([]int, 0),

		tokenPool:    token2.NewPool(position2.DefaultBlockSize),
		positionPool: position2.NewPool(token2.DefaultBlockSize),
		newLines:     NewLines{make([]int, 0, 128)},
	}

	initLexer(lex)

	return lex
}

func (lex *Lexer) setTokenPos(token *token2.Token) {
	pos := lex.positionPool.Get()

	pos.StartLine = lex.newLines.GetLine(lex.ts)
	pos.EndLine = lex.newLines.GetLine(lex.te - 1)
	pos.StartPos = lex.ts
	pos.EndPos = lex.te

	token.Pos = pos
}

func (lex *Lexer) addFreeFloatingToken(t *token2.Token, id token2.ID, ps, pe int) {
	skippedTkn := lex.tokenPool.Get()
	skippedTkn.ID = id
	skippedTkn.Value = lex.data[ps:pe]

	lex.setTokenPos(skippedTkn)

	if t.FreeFloating == nil {
		t.FreeFloating = make([]*token2.Token, 0, 2)
	}

	t.FreeFloating = append(t.FreeFloating, skippedTkn)
}

func (lex *Lexer) growCallStack() {
	if lex.top == len(lex.stack) {
		lex.stack = append(lex.stack, 0)
	}
}

func (lex *Lexer) call(state int, fnext int) {
	lex.growCallStack()

	lex.stack[lex.top] = state
	lex.top++

	lex.p++
	lex.cs = fnext
}

func (lex *Lexer) ret(n int) {
	lex.top = lex.top - n
	if lex.top < 0 {
		lex.top = 0
	}
	lex.cs = lex.stack[lex.top]
	lex.p++
}

func (lex *Lexer) ungetStr(s string) {
	tokenStr := string(lex.data[lex.ts:lex.te])
	if strings.HasSuffix(tokenStr, s) {
		lex.ungetCnt(len(s))
	}
}

func (lex *Lexer) ungetCnt(n int) {
	lex.p = lex.p - n
	lex.te = lex.te - n
}

func (lex *Lexer) error(msg string) {
	if lex.errHandlerFunc == nil {
		return
	}

	pos := position2.NewPos(
		lex.newLines.GetLine(lex.ts),
		lex.newLines.GetLine(lex.te-1),
		lex.ts,
		lex.te,
	)

	lex.errHandlerFunc(errors.NewError(msg, pos))
}

func (lex *Lexer) isNotNewLine() bool {
	if lex.data[lex.p] == '\n' && lex.data[lex.p-1] == '\r' {
		return true
	}

	return lex.data[lex.p-1] != '\n' && lex.data[lex.p-1] != '\r'
}
