package representator

import (
	"fmt"

	"github.com/i582/CodeQuery/pkg/models"
)

func PrintGraph(db *models.Database, graph *models.Graph) {
	fmt.Println("Function " + db.GetFuncByID(graph.Root).FullName() + " dependencies")
	fmt.Println()
	printNode(db, graph, graph.Root, 0, map[int64]struct{}{})
}

func printNode(db *models.Database, graph *models.Graph, id int64, level int64, visited map[int64]struct{}) {
	if _, ok := visited[id]; ok {
		return
	}
	visited[id] = struct{}{}

	var node *models.Node
	for _, graph := range db.Graphs {
		nod, ok := graph.Nodes[id]
		if ok {
			node = nod
			break
		}
	}

	for i := int64(0); i < level; i++ {
		fmt.Print("   ")
	}

	fmt.Print("└─> ")
	fmt.Println(node.Data.Fqn.Name_)

	ids := graph.Graph[id]
	for _, id := range ids {
		printNode(db, graph, id, level+1, visited)
	}
}
