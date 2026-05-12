package main

import (
	"github.com/aungshanbo/a9r/configs"
	"github.com/aungshanbo/a9r/models"
	"github.com/aungshanbo/a9r/ui"
	"github.com/rivo/tview"
)

func main() {

	profiles := configs.Loadconfig()
	regions := configs.Getregion()

	app := tview.NewApplication()

	state := &models.AppState{}

	var selectedProfile string
	var selectedRegion string

	autoRefresh := false

	table := ui.NewTable()

	statusBar := tview.NewTextView()

	statusBar.SetBorder(true)
	statusBar.SetTitle("Help")

	statusBar.SetText(
		"TAB=switch | /=search | r=refresh | a=auto | q=quit",
	)

	leftPanel,
		form,
		statusView := ui.NewForm(
		app,
		table,
		profiles,
		regions,
		&selectedProfile,
		&selectedRegion,
	)

	layout := ui.NewLayout(
		leftPanel,
		table,
		statusBar,
	)

	table.SetSelectedFunc(func(row, column int) {

		if row == 0 {
			return
		}

		ui.ShowInstanceModal(
			app,
			layout,
			table,
			row,
		)
	})

	ui.BindTableKeys(
		app,
		form,
		table,
	)

	ui.BindGlobalKeys(
		app,
		table,
		statusBar,
		statusView,
		state,
		&selectedProfile,
		&selectedRegion,
		&autoRefresh,
	)

	ui.StartAutoRefresh(
		app,
		table,
		statusView,
		state,
		&selectedProfile,
		&selectedRegion,
		&autoRefresh,
	)

	app.SetFocus(form)

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
