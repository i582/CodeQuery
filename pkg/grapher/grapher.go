package grapher

import (
	"github.com/i582/CodeQuery/pkg/models"
)

type Grapher struct {
	Database *models.Database
	Graph    *models.Graph
}

func NewGrapher(database *models.Database) *Grapher {
	return &Grapher{Database: database}
}
