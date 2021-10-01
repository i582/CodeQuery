package representator

import (
	"github.com/alexeyco/simpletable"
	"github.com/gookit/color"
	"github.com/i582/CodeQuery/pkg/models"
)

func GetTableFunctionsCallsRepr(f *models.FuncCallTable) string {
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

	for _, call := range f.Data {
		var r []*simpletable.Cell

		if f.NeedShowCol("args") {
			var cellValue string
			for _, arg := range call.Args_ {
				cellValue += splitText(arg.Value) + "\n"
			}

			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: cellValue})
		}
		if f.NeedShowCol("args_count") {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignRight, Text: color.Gray.Sprint(len(call.Args_))})
		}
		if f.NeedShowCol("file") {
			r = append(r, &simpletable.Cell{Text: splitText(call.Pos.String())})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	return table.String()
}
