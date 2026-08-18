package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowLoading(
	app *tview.Application,
	pages *tview.Pages,
	table *tview.Table,
	title string,
) {

	loading := tview.NewTextView()

	loading.
		SetTextAlign(tview.AlignCenter).
		SetText("Loading " + title + "...")

	loading.SetBorder(true)
	loading.SetTitle("A9R")

	loading.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {
			return nil
		},
	)

	pages.AddPage(
		"loading",
		loading,
		true,
		false,
	)

	pages.SwitchToPage("loading")

	app.SetFocus(loading)
}
