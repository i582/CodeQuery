package symbols

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/meta"
)

type FunctionType int64

const (
	MainFunc FunctionType = iota
	LocalFunc
	ExternFunc
)

type FunctionArg struct {
	Expr ir.Node
}

type FunctionCall struct {
	Function *Function
	Position meta.ElementPosition
	Args     []FunctionArg
}

type FunctionCalls struct {
	mtx   sync.Mutex
	Calls map[meta.ElementPosition]FunctionCall
}

func NewFunctionCalls() *FunctionCalls {
	return &FunctionCalls{Calls: map[meta.ElementPosition]FunctionCall{}}
}

func (c *FunctionCalls) Add(call FunctionCall) {
	c.mtx.Lock()
	c.Calls[call.Position] = call
	c.mtx.Unlock()
}

// Function is a structure for storing information about a function.
type Function struct {
	Name       string
	Type       FunctionType
	Pos        meta.ElementPosition
	NeedRemove bool

	Globals *GlobalVars

	Calls *FunctionCalls

	FunctionsCalls *FunctionCalls

	Called   *Functions
	CalledBy *Functions
}

// HumanReadableName returns a string with a name that is understandable.
func (f *Function) HumanReadableName() string {
	if f.Type == MainFunc {
		name := f.Name
		path := name[strings.LastIndex(name, "$")+1:]

		wd, err := os.Getwd()
		if err == nil {
			relPath, err := filepath.Rel(wd, path)
			if err == nil {
				path = relPath
			}
		}

		path = filepath.ToSlash(path)

		return fmt.Sprintf("file '%s' scope", path)
	}

	return strings.TrimPrefix(f.Name, `\`)
}

func (f *Function) String() string {
	return f.Name
}

type Functions struct {
	mtx       sync.Mutex
	Functions map[string]*Function
}

func NewFunctions() *Functions {
	return &Functions{Functions: map[string]*Function{}}
}

func (f *Functions) Get(name string) (*Function, bool) {
	fun, ok := f.Functions[name]
	return fun, ok
}

func (f *Functions) Raw() map[string]*Function {
	return f.Functions
}

func (f *Functions) Len() int {
	return len(f.Functions)
}

func (f *Functions) Add(fun *Function) {
	f.mtx.Lock()
	f.Functions[fun.Name] = fun
	f.mtx.Unlock()
}
