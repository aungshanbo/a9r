package ui

import (
	"github.com/aungshanbo/a9r/models"
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
				SetExpansion(1),
		)
	}

	// =========================
	// ROWS
	// =========================
	for r, row := range rows {

		for c, col := range row {

			table.SetCell(
				r+1,
				c,
				tview.NewTableCell(col).
					SetExpansion(1),
			)
		}
	}
}
