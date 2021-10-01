package preprocess

import (
	"github.com/i582/CodeQuery/pkg/models"
)

// NodesToGraphs splits the graph into connectivity components.
//
// Graphs from one node are skipped.
func NodesToGraphs(nodes models.Nodes) []*models.Graph {
	visited := make(map[int64]struct{}, len(nodes))
	graphs := make([]*models.Graph, 0, 10)
	queue := make([]*models.Node, 0, 10)
	graphFunctions := make(models.Nodes, 10)

	for _, node := range nodes {
		if _, ok := visited[node.ID]; ok {
			continue
		}

		queue = append(queue, node)

		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]

			if _, ok := visited[node.ID]; ok {
				continue
			}
			visited[node.ID] = struct{}{}

			graphFunctions[node.ID] = node

			for _, index := range node.Next {
				queue = append(queue, nodes[index])
			}

			for _, index := range node.Prev {
				queue = append(queue, nodes[index])
			}
		}

		if len(graphFunctions) > 1 {
			graph, revGraph := functionsToRawGraphs(graphFunctions)

			graphs = append(graphs, &models.Graph{
				Nodes:    graphFunctions,
				Graph:    graph,
				RevGraph: revGraph,
				Root:     -1,
			})
		}

		graphFunctions = make(models.Nodes, 10)
	}

	return graphs
}

func functionsToRawGraphs(funcs models.Nodes) (graph, revGraph models.RawGraph) {
	graph = make(models.RawGraph, len(funcs))
	revGraph = make(models.RawGraph, len(funcs))

	for _, node := range funcs {
		graph[node.ID] = node.Next
		revGraph[node.ID] = node.Prev
	}

	return graph, revGraph
}
