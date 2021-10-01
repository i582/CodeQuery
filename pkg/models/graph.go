package models

type Nodes map[int64]*Node

// NodesIDs is an alias for a slice of graph nodes for which a Remove
// function is defined for ease of interaction.
type NodesIDs []int64

// RawGraph is a map that stores all the functions that are called
// in the function or the function in which the function is called.
type RawGraph map[int64]NodesIDs

// Graph is a structure for storing the complete call graph, as
// well as the functions that are included in it.
type Graph struct {
	Nodes Nodes

	Root     int64
	Graph    RawGraph
	RevGraph RawGraph
}

// Remove is a function that removes the passed node from the Graph.
//
// Note: Node is not removed from the RevGraph.
func (g *Graph) Remove(node *Node) {
	callers := g.RevGraph[node.ID]
	for _, caller := range callers {
		g.Graph[caller] = g.Graph[caller].Remove(node)
	}
}

// Node is a structure for storing node.
type Node struct {
	ID   int64
	Data Func

	Next NodesIDs
	Prev NodesIDs
}

// NewNode creates new node with ID.
func NewNode(id int64) *Node {
	return &Node{
		ID: id,
	}
}

// String method for debugging.
func (n *Node) String() string {
	return n.Data.Fqn.String()
}

// Remove is a function that removes the passed node from the slice if any.
func (n Nodes) Remove(node *Node) Nodes {
	delete(n, node.ID)
	return n
}

// Remove is a function that removes the passed node from the slice if any.
func (n NodesIDs) Remove(node *Node) NodesIDs {
	removeHelper := func(s NodesIDs, i int) NodesIDs {
		s[i] = s[len(s)-1]
		return s[:len(s)-1]
	}

	index := -1
	for i, fun := range n {
		if fun == node.ID {
			index = i
			continue
		}
	}

	if index != -1 {
		return removeHelper(n, index)
	}

	return n
}
