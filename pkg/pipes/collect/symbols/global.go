package symbols

import (
	"fmt"
	"sync"
)

type GlobalVar struct {
	Name     string
	UseCount int64
}

type GlobalVars struct {
	mtx        sync.Mutex
	GlobalVars map[string]*GlobalVar
}

func NewGlobals() *GlobalVars {
	return &GlobalVars{
		GlobalVars: map[string]*GlobalVar{},
	}
}

func (g *GlobalVars) Get(name string) (*GlobalVar, bool) {
	g.mtx.Lock()
	vr, ok := g.GlobalVars[name]
	g.mtx.Unlock()
	return vr, ok
}

func (g *GlobalVars) Add(variable *GlobalVar) {
	if g == nil {
		fmt.Println()
	}
	g.mtx.Lock()
	g.GlobalVars[variable.Name] = variable
	g.mtx.Unlock()
}
