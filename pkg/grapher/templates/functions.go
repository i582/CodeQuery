package templates

import (
	"github.com/i582/CodeQuery/pkg/graph"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/utils"
)

func TemplateFunctionNode(c *models.Func) *graph.Node {
	name := utils.NameToIdentifier(c.FullName())

	label := utils.NormalizeSlashes(c.FullName())

	return &graph.Node{
		Name: name,
		Styles: graph.NodeStyles{
			Label:     label,
			Shape:     "rect",
			FillColor: DefaultFillColor,
			EdgeColor: DefaultOutlineColor,
			Style:     "filled",
			FontSize:  12,
		},
	}
}
