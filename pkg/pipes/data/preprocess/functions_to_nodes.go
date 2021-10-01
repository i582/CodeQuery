package preprocess

import (
	"strings"

	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/pipes/collect/symbols"
)

// FunctionsToNodes is a function that converts passed functions into
// a set of nodes that represent those functions.
func FunctionsToNodes(funcs *symbols.Functions, db *models.Database) models.Nodes {
	nodes := make(models.Nodes, funcs.Len())
	visited := make(map[*symbols.Function]*models.Node, funcs.Len())

	for _, fun := range funcs.Functions {
		if strings.Contains(fun.Name, "rc$") {
			continue
		}
		functionToNode(nodes, fun, db, visited)
	}

	return nodes
}

func functionToNode(nodes models.Nodes, fun *symbols.Function, db *models.Database, visited map[*symbols.Function]*models.Node) *models.Node {
	if node, ok := visited[fun]; ok {
		return node
	}

	convertedFun := db.Funcs[convertName(fun.Name)]
	if convertedFun == nil {
		return &models.Node{}
	}

	node := models.NewNode(convertedFun.ID)
	nodes[node.ID] = node

	visited[fun] = node

	node.Data = *convertedFun

	for _, called := range fun.Called.Functions {
		nextNode := functionToNode(nodes, called, db, visited)
		node.Next = append(node.Next, nextNode.ID)
	}

	for _, calledBy := range fun.CalledBy.Functions {
		prevNode := functionToNode(nodes, calledBy, db, visited)
		node.Prev = append(node.Prev, prevNode.ID)
	}

	return node
}
