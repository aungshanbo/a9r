package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewForm(
	app *tview.Application,
	table *tview.Table,
	profiles []string,
	regions []string,
	selectedProfile *string,
	selectedRegion *string,
) (
	tview.Primitive,
	*tview.Form,
	*tview.TextView,
) {

	// ==========================================
	// GLOBAL STYLE
	// ==========================================
	tview.Styles.PrimaryTextColor = tcell.ColorWhite
	tview.Styles.ContrastBackgroundColor = tcell.ColorBlack

	// ==========================================
	// FORM
	// ==========================================
	form := tview.NewForm()

	form.SetFieldTextColor(tcell.ColorWhite)
	form.SetFieldBackgroundColor(tcell.ColorBlack)

	form.SetLabelColor(tcell.ColorWhite)

	form.SetBorderColor(tcell.ColorWhite)
	form.SetTitleColor(tcell.ColorWhite)

	form.SetButtonBackgroundColor(tcell.ColorBlack)
	form.SetButtonTextColor(tcell.ColorWhite)

	dropdownOpen := false

	// ==========================================
	// PROFILE
	// ==========================================
	form.AddDropDown(
		"Profile : ",
		profiles,
		0,
		func(option string, index int) {
			*selectedProfile = option
		},
	)

	// ==========================================
	// REGION
	// ==========================================
	form.AddDropDown(
		"Region  : ",
		regions,
		0,
		func(option string, index int) {
			*selectedRegion = option
		},
	)

	// ==========================================
	// STATUS VIEW
	// ==========================================
	statusView := tview.NewTextView()

	statusView.SetDynamicColors(true)
	statusView.SetText("Status : Ready")

	// ==========================================
	// FORM STYLE
	// ==========================================
	form.SetBorder(true)
	form.SetTitle("A9R")

	// ==========================================
	// KEYBOARD
	// ==========================================
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		index, _ := form.GetFocusedItemIndex()

		switch event.Key() {

		case tcell.KeyEnter:

			if index <= 1 {
				dropdownOpen = !dropdownOpen
			}

			return event

		case tcell.KeyDown:

			if dropdownOpen {
				return event
			}

			if index < 1 {
				form.SetFocus(index + 1)
			}

			return nil

		case tcell.KeyUp:

			if dropdownOpen {
				return event
			}

			if index > 0 {
				form.SetFocus(index - 1)
			}

			return nil

		case tcell.KeyTAB:

			dropdownOpen = false

			app.SetFocus(table)

			return nil
		}

		return event
	})

	// ==========================================
	// LEFT PANEL
	// ==========================================
	leftPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 7, 1, true).
		AddItem(statusView, 1, 1, false)

	return leftPanel, form, statusView
}
