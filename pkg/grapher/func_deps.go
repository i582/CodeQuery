package grapher

import (
	"github.com/i582/CodeQuery/pkg/graph"
	"github.com/i582/CodeQuery/pkg/grapher/templates"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/utils"
)

func (g *Grapher) FunctionDeps(funcGraph *models.Graph) string {
	g.Graph = funcGraph
	root := g.Database.GetFuncByID(funcGraph.Root)

	graphName := "GraphFor_" + utils.NameToIdentifier(root.Name())
	functionGraph := &graph.Graph{
		Name:       graphName,
		IsSubgraph: false,
		GraphStyle: graph.Styles{
			Label:      "Function " + root.FullName() + " dependencies",
			Padding:    2.0,
			NodeMargin: 1.5,
		},
		NodeStyle: graph.NodeStyles{},
		EdgeStyle: templates.TemplateFunctionConnectionEdgeStyle(),
	}

	g.funcDepsRecursive(functionGraph, root, 0, 10)

	funcNode, found := functionGraph.GetNodeInSubgraphs(utils.NameToIdentifier(root.FullName()))
	if found {
		funcNode.Styles.FillColor = templates.FillColorLevel3
		funcNode.Styles.EdgeColor = templates.OutlineColorLevel3
		funcNode.Scale(1.7)
	}

	return functionGraph.String()
}

func (g *Grapher) funcDepsRecursive(functionGraph *graph.Graph, f *models.Func, levelRecursion, maxRecursion int64) {
	mainFunctionSubGraph := g.createSubGraphForFunctionClass(f, functionGraph)
	mainFuncNode, _ := mainFunctionSubGraph.AddNode(templates.TemplateFunctionNode(f))

	if levelRecursion > maxRecursion {
		return
	}

	for _, id := range g.Graph.Graph[f.ID] {
		fun := g.Database.GetFuncByID(id)
		subGraph := g.createSubGraphForFunctionClass(fun, functionGraph)
		subGraph.AddNode(templates.TemplateFunctionNode(fun))

		g.funcDepsRecursive(functionGraph, fun, levelRecursion+1, maxRecursion)
	}

	for _, id := range g.Graph.Graph[f.ID] {
		fun := g.Database.GetFuncByID(id)
		funcNode, found := functionGraph.GetNodeInSubgraphs(utils.NameToIdentifier(fun.FullName()))
		if !found {
			continue
		}

		functionGraph.AddEdgeByNode(mainFuncNode, funcNode, graph.EdgeStyles{Color: templates.OutlineColorLevel2})
	}
}

func (g *Grapher) createSubGraphForFunctionClass(function *models.Func, functionGraph *graph.Graph) *graph.Graph {
	return g.createSubGraphForClass(function.ClassName(), functionGraph)
}

func (g *Grapher) createSubGraphForClass(class string, functionGraph *graph.Graph) *graph.Graph {
	var subGraph *graph.Graph
	var found bool

	if class != "" {
		subGraphName := utils.NameToIdentifier(class)
		subGraph, found = functionGraph.GetSubGraph(subGraphName)
		if !found {
			subGraph = functionGraph.AddSubGraph(&graph.Graph{
				Name:       subGraphName,
				IsSubgraph: false,
				GraphStyle: graph.Styles{
					Label:       utils.NormalizeSlashes(class),
					BorderColor: templates.DefaultOutlineColor,
					FontColor:   templates.DefaultOutlineColor,
				},
			})
		}
	} else {
		subGraph, found = functionGraph.GetSubGraph("globalScope")
		if !found {
			subGraph = functionGraph.AddSubGraph(&graph.Graph{
				Name:       "globalScope",
				IsSubgraph: false,
				GraphStyle: graph.Styles{
					Label:       "Global Scope",
					BorderColor: templates.DefaultOutlineColor,
					FontColor:   templates.DefaultOutlineColor,
				},
			})
		}
	}
	return subGraph
}
