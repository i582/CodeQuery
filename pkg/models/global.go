package models

import (
	"github.com/i582/CodeQuery/pkg/querylang/models"
)

type Global struct {
	ID int64

	Name_    string
	UseCount int64
}

func (g Global) Id() int64 {
	return g.ID
}

func (g Global) Name() string {
	return g.Name_
}

func (g Global) CountUse() int64 {
	return g.UseCount
}

type Globals struct {
	ByID   map[int64]Global
	ByName map[string]Global
}

func (g Globals) Count() int64 {
	return int64(len(g.ByID))
}

func (g Globals) Iterate(cb func(global models.IGlobal) bool) {
	for _, global := range g.ByID {
		if !cb(global) {
			break
		}
	}
}

func (g Globals) Has(global models.IGlobal) bool {
	_, ok := g.ByID[global.Id()]
	return ok
}

func (g Globals) Contains(globals models.IGlobals) bool {
	if g.Count() == 0 {
		return false
	}

	if g.Count() < globals.Count() {
		for _, global := range g.ByID {
			if globals.Has(global) {
				return true
			}
		}

		return false
	}

	contains := false
	globals.Iterate(func(global models.IGlobal) bool {
		if _, ok := g.ByID[global.Id()]; ok {
			contains = true
			return false
		}
		return true
	})
	return contains
}

type GlobalTable struct {
	Cols []TableCol
	Data []Global
}

func (t *GlobalTable) Iterate(f func(fun models.IGlobal)) {
	for _, d := range t.Data {
		f(d)
	}
}

func (t *GlobalTable) Length() int64 {
	return int64(len(t.Data))
}

func (t *GlobalTable) NeedShowCol(name string) bool {
	for _, col := range t.Cols {
		if col.IdName == name {
			return true
		}
	}
	return false
}
