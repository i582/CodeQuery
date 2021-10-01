package representator

import (
	"github.com/alexeyco/simpletable"
	"github.com/gookit/color"
	"github.com/i582/CodeQuery/pkg/models"
)

func GetTableGlobalsRepr(f *models.GlobalTable) string {
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

	for _, global := range f.Data {
		name := global.Name()

		var r []*simpletable.Cell

		if f.NeedShowCol("id") {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: color.Gray.Sprint(global.ID)})
		}
		if f.NeedShowCol("name") {
			r = append(r, &simpletable.Cell{Text: name})
		}
		if f.NeedShowCol("uses") {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: ColorOutputIntZeroableValue(global.UseCount)})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	return table.String()
}
