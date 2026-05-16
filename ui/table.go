package ui

import (
	"github.com/aungshanbo/a9r/models"
	"github.com/aungshanbo/a9r/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewTable() *tview.Table {

	table := tview.NewTable()

	table.SetBorders(true)
	table.SetBorder(true)
	table.SetSelectable(true, false)
	table.SetFixed(1, 0)

	table.SetSelectedStyle(
		tcell.StyleDefault.
			Background(tcell.ColorBlue).
			Foreground(tcell.ColorWhite),
	)

	return table
}

func DrawTable(
	table *tview.Table,
	headers []models.TableColumn,
	rows [][]string,
) {

	table.Clear()

	// =========================
	// HEADER
	// =========================
	for i, h := range headers {

		table.SetCell(
			0,
			i,
			tview.NewTableCell(h.Title).
				SetAttributes(tcell.AttrBold).
				SetSelectable(false).
				SetExpansion(h.Expansion),
		)
	}

	// =========================
	// ROWS
	// =========================
	for r, row := range rows {

		for c, col := range row {

			cell := tview.NewTableCell(col).
				SetExpansion(headers[c].Expansion)

			// state column color
			if headers[c].Title == "State" {

				cell.SetTextColor(
					utils.Getstatecolor(col),
				)
			}

			table.SetCell(
				r+1,
				c,
				cell,
			)
		}
	}
}
