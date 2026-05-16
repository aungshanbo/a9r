package ui

import (
	"fmt"
	"time"

	"github.com/aungshanbo/a9r/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BindTableKeys(
	table *tview.Table,
) {

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch {

		// allow "/" to global handler
		case event.Rune() == '/':

			return event

		// move down
		case event.Rune() == 'j',
			event.Key() == tcell.KeyDown:

			moveTableDown(table)
			return nil

		// move up
		case event.Rune() == 'k',
			event.Key() == tcell.KeyUp:

			moveTableUp(table)
			return nil
		}

		return event
	})
}

func BindGlobalKeys(
	app *tview.Application,
	table *tview.Table,
	profileDropDown *tview.DropDown,
	regionDropDown *tview.DropDown,
	resourceDropDown *tview.DropDown,
	statusBar *tview.TextView,
	statusView *tview.TextView,
	state *models.AppState,
	selectedProfile *string,
	selectedRegion *string,
	selectedResource *string,
	autoRefresh *bool,
) {
	searchMode := false
	var searchTimer *time.Timer
	// ==========================================
	// STATUS BAR
	// ==========================================
	updateStatusBar := func() {

		statusBar.SetDynamicColors(true)

		left := "TAB=switch | /=search | r=refresh | a=auto | q=quit"

		if *autoRefresh {
			left = "Auto: ON | TAB=switch | /=search | r=refresh | q=quit"
		}

		if searchMode {
			left = "Search Mode | ESC=clear"
		}

		if state.Filter == "" {
			statusBar.SetText(left)
			return
		}

		right := fmt.Sprintf("Search: %s", state.Filter)

		_, _, width, _ := statusBar.GetInnerRect()

		padding := width - len(left) - len(right)

		if padding < 1 {
			padding = 1
		}

		statusBar.SetText(
			left +
				fmt.Sprintf("%*s", padding, "") +
				right,
		)
	}

	// ==========================================
	// DRAW FILTERED TABLE
	// ==========================================
	drawFilteredTable := func() {

		if state.CurrentResource == nil {
			return
		}

		filtered := FilterRows(
			state.CurrentResource.Rows,
			state.Filter,
		)

		DrawTable(
			table,
			state.CurrentResource.Headers,
			filtered,
		)

		if len(filtered) > 0 {
			table.Select(1, 0)
		}
	}

	// ==========================================
	// EXIT SEARCH MODE
	// ==========================================
	exitSearchMode := func(clear bool) {

		searchMode = false

		if clear {
			state.Filter = ""
		}

		updateStatusBar()

		// immediate refresh
		if *selectedProfile != "" &&
			*selectedRegion != "" {

			go RefreshTable(
				app,
				table,
				statusView,
				state,
				*selectedProfile,
				*selectedRegion,
				*selectedResource,
			)
		}
	}

	updateStatusBar()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// ==========================================
		// TAB switch
		// ==========================================

		focusables := []tview.Primitive{
			profileDropDown,
			regionDropDown,
			resourceDropDown,
			table,
		}

		if event.Key() == tcell.KeyTAB {

			current := app.GetFocus()

			for i, item := range focusables {

				if current == item {

					next := (i + 1) % len(focusables)

					app.SetFocus(focusables[next])

					return nil
				}
			}
		}

		// ==========================================
		// SEARCH MODE
		// ==========================================
		if searchMode {

			switch {
			// clear + exit
			case event.Key() == tcell.KeyEsc:

				exitSearchMode(true)
				return nil

			// move down
			case event.Rune() == 'j',
				event.Key() == tcell.KeyDown:

				moveTableDown(table)
				return nil

			// move up
			case event.Rune() == 'k',
				event.Key() == tcell.KeyUp:

				moveTableUp(table)
				return nil

			// backspace
			case event.Key() == tcell.KeyBackspace,
				event.Key() == tcell.KeyBackspace2:

				if len(state.Filter) > 0 {
					state.Filter = state.Filter[:len(state.Filter)-1]
				}

				if searchTimer != nil {
					searchTimer.Stop()
				}

				searchTimer = time.AfterFunc(
					120*time.Millisecond,
					func() {
						app.QueueUpdateDraw(func() {
							drawFilteredTable()
						})
					},
				)
				updateStatusBar()

				return nil
			}

			// typing
			if event.Rune() != 0 {

				state.Filter += string(event.Rune())

				updateStatusBar()

				if searchTimer != nil {
					searchTimer.Stop()
				}

				searchTimer = time.AfterFunc(
					120*time.Millisecond,
					func() {
						app.QueueUpdateDraw(func() {
							drawFilteredTable()
						})
					},
				)

				return nil
			}

			return nil
		}

		// ==========================================
		// NORMAL MODE
		// ==========================================
		switch {

		// quit
		case event.Rune() == 'q':

			app.Stop()
			return nil

		// refresh
		case event.Rune() == 'r':

			if *selectedProfile != "" &&
				*selectedRegion != "" {

				go RefreshTable(
					app,
					table,
					statusView,
					state,
					*selectedProfile,
					*selectedRegion,
					*selectedResource,
				)
			}

			return nil

		// search mode
		case event.Rune() == '/':

			if app.GetFocus() != table {
				return event
			}

			searchMode = true
			state.Filter = ""

			updateStatusBar()

			return nil

		// auto refresh
		case event.Rune() == 'a':

			*autoRefresh = !*autoRefresh

			updateStatusBar()

			// instant refresh
			if *autoRefresh &&
				*selectedProfile != "" &&
				*selectedRegion != "" {

				go RefreshTable(
					app,
					table,
					statusView,
					state,
					*selectedProfile,
					*selectedRegion,
					*selectedResource,
				)
			}

			return nil
		}

		return event
	})
}

func StartAutoRefresh(
	app *tview.Application,
	table *tview.Table,
	statusView *tview.TextView,
	state *models.AppState,
	selectedProfile *string,
	selectedRegion *string,
	selectedResource *string,
	autoRefresh *bool,
) {

	go func() {

		for {

			time.Sleep(10 * time.Second)

			if *autoRefresh &&
				*selectedProfile != "" &&
				*selectedRegion != "" {

				RefreshTable(
					app,
					table,
					statusView,
					state,
					*selectedProfile,
					*selectedRegion,
					*selectedResource,
				)
			}
		}
	}()
}

func moveTableDown(table *tview.Table) {

	r, c := table.GetSelection()

	maxRow := table.GetRowCount() - 1

	if maxRow < 1 {
		return
	}

	if r < maxRow {
		table.Select(r+1, c)
	}
}

func moveTableUp(table *tview.Table) {

	r, c := table.GetSelection()

	if r > 1 {
		table.Select(r-1, c)
	}
}
