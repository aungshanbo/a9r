package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func ShowInstanceModal(
	app *tview.Application,
	layout tview.Primitive,
	table *tview.Table,
	row int,
) {

	id := table.GetCell(row, 0).Text
	name := table.GetCell(row, 1).Text
	state := table.GetCell(row, 2).Text
	itype := table.GetCell(row, 3).Text
	az := table.GetCell(row, 4).Text
	privateIP := table.GetCell(row, 5).Text
	publicIP := table.GetCell(row, 6).Text

	modal := tview.NewModal().
		SetText(fmt.Sprintf(
			"ID: %s\nName: %s\nState: %s\nType: %s\nAZ: %s\nPrivateIP: %s\nPublicIP: %s",
			id,
			name,
			state,
			itype,
			az,
			privateIP,
			publicIP,
		)).
		AddButtons([]string{"Close"}).
		SetDoneFunc(func(i int, l string) {

			app.SetRoot(layout, true)
			app.SetFocus(table)
		})

	app.SetRoot(modal, true)
}