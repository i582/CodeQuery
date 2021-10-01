package models

func (g *Graph) FindPaths(db *Database, parentID, childID int64, maxLevel int64) [][]*Func {
	return g.findPaths(db, db.GetFuncByID(parentID), db.GetFuncByID(childID), nil, map[int64]struct{}{}, 0, maxLevel)
}

func (g *Graph) ReversePaths(funcs [][]*Func) {
	for _, path := range funcs {
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
	}
}

func (g *Graph) findPaths(db *Database, parent, child *Func, callstack []*Func, visited map[int64]struct{}, currentLevel, maxLevel int64) [][]*Func {
	if len(g.RevGraph[parent.ID]) == 0 {
		return nil
	}

	if callstack == nil {
		callstack = []*Func{parent}
	}

	if parent == child {
		return [][]*Func{callstack}
	}

	if currentLevel >= maxLevel {
		return [][]*Func{}
	}

	var callstacks [][]*Func

	for _, calledID := range g.RevGraph[parent.ID] {
		called := db.GetFuncByID(calledID)
		newCallstack := copyCallstack(callstack)
		newVisited := copyVisited(visited)

		newCallstack = append(newCallstack, called)

		if _, ok := newVisited[calledID]; ok {
			continue
		}
		if called.ID == parent.ID {
			continue
		}

		newVisited[called.ID] = struct{}{}

		if called.ID == child.ID {
			callstacks = append(callstacks, newCallstack)
			continue
		}

		foundCallstack := g.findPaths(db, called, child, newCallstack, newVisited, currentLevel+1, maxLevel)
		if len(foundCallstack) != 0 {
			callstacks = append(callstacks, foundCallstack...)
		}
	}

	return callstacks
}

func copyCallstack(callstack []*Func) []*Func {
	tmp := make([]*Func, len(callstack))
	copy(tmp, callstack)
	return tmp
}

func copyVisited(visited map[int64]struct{}) map[int64]struct{} {
	targetMap := make(map[int64]struct{})

	for key, value := range visited {
		targetMap[key] = value
	}

	return targetMap
}
