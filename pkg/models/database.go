package models

type Database struct {
	Graphs  []*Graph
	Funcs   FuncMap
	Globals Globals
}

func (d *Database) GetFuncByID(id int64) *Func {
	var fun Func
	for _, graph := range d.Graphs {
		node, ok := graph.Nodes[id]
		if ok {
			fun = node.Data
		}
	}
	return &fun
}
