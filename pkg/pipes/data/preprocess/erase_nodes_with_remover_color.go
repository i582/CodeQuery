package preprocess

import (
	"github.com/i582/CodeQuery/pkg/models"
)

// EraseNodesWithRemoverColor is function for drop '@color remover' functions from a
// call graph.
//
// To perform calculating NextWithColors and checking color rules from the palette,
// we need to drop '@color remover' functions from a call graph completely, like
// they don't exist at all, this special color is for manual cutting connectivity
// rules, allowing to explicitly separate recursively-joint components
func EraseNodesWithRemoverColor(graph *models.Graph) {
	for _, fun := range graph.Nodes {
		if fun.Data.Rem {
			graph.Remove(fun)
			graph.Nodes.Remove(fun)
		}
	}
}
