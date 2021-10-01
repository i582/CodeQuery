package grapher

import (
	"github.com/i582/CodeQuery/pkg/graph"
	"github.com/i582/CodeQuery/pkg/grapher/templates"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/utils"
)

func (g *Grapher) Paths(paths [][]*models.Func) string {
	graphName := "GraphFor_"
	pathsGraph := &graph.Graph{
		Name:       graphName,
		IsSubgraph: false,
		GraphStyle: graph.Styles{
			Label:      "Functions paths dependency",
			Padding:    2.0,
			NodeMargin: 1.5,
		},
		NodeStyle: graph.NodeStyles{},
		EdgeStyle: templates.TemplateFunctionConnectionEdgeStyle(),
	}

	for _, path := range paths {
		for _, node := range path {
			pathNode := templates.TemplateFunctionNode(node)
			pathNode, _ = pathsGraph.AddNode(pathNode)
		}

		for i := 0; i < len(path)-1; i++ {
			node, _ := pathsGraph.GetNode(utils.NameToIdentifier(path[i].FullName()))
			nextNode, _ := pathsGraph.GetNode(utils.NameToIdentifier(path[i+1].FullName()))

			pathsGraph.AddEdgeByNode(node, nextNode, templates.TemplateFunctionConnectionEdgeStyle())
		}
	}

	return pathsGraph.String()
}
