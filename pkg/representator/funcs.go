package representator

import (
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/gookit/color"
	"github.com/i582/CodeQuery/pkg/models"
)

func GetTableFunctionsRepr(f *models.FuncTable) string {
	if f == nil {
		return ""
	}

	table := simpletable.New()
	table.SetStyle(simpletable.StyleCompactLite)
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{},
	}

	for _, col := range f.Cols {
		table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  color.Green.Sprint(col.Name) + "\n(" + col.IdName + ")",
		})
	}

	for _, fun := range f.Data {
		if strings.Contains(fun.Fqn.Name_, "src$") {
			continue
		}
		name := fun.FullName()

		if len(f.Cols) > 3 {
			name = splitText(fun.FullName())
		}

		var r []*simpletable.Cell

		if f.NeedShowCol("id") {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: color.Gray.Sprint(fun.ID)})
		}
		if f.NeedShowCol("name") {
			r = append(r, &simpletable.Cell{Text: name})
		}
		if f.NeedShowCol("uses") {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: ColorOutputIntZeroableValue(fun.UseCount)})
		}
		if f.NeedShowCol("globals") {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: ColorOutputIntZeroableValue(fun.Globals().Count())})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	return table.String()
}
