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
	selectedResource *string,
) (
	tview.Primitive,
	*tview.DropDown,
	*tview.DropDown,
	*tview.DropDown,
	*tview.TextView,
) {

	// ==================================================
	// PROFILE
	// ==================================================

	profileDropDown := tview.NewDropDown()

	profileDropDown.SetLabel("Profile : ")
	profileDropDown.SetOptions(
		profiles,
		func(option string, index int) {
			*selectedProfile = option
		},
	)

	profileDropDown.SetCurrentOption(0)

	if len(profiles) > 0 {
		*selectedProfile = profiles[0]
	}

	// ==================================================
	// REGION
	// ==================================================

	regionDropDown := tview.NewDropDown()

	regionDropDown.SetLabel("Region  : ")
	regionDropDown.SetOptions(
		regions,
		func(option string, index int) {
			*selectedRegion = option
		},
	)

	regionDropDown.SetCurrentOption(0)

	if len(regions) > 0 {
		*selectedRegion = regions[0]
	}

	// ==================================================
	// RESOURCE
	// ==================================================

	resourceOptions := []string{
		"EC2",
		"S3",
	}

	resourceDropDown := tview.NewDropDown()

	resourceDropDown.SetLabel("Resource: ")
	resourceDropDown.SetOptions(
		resourceOptions,
		func(option string, index int) {
			*selectedResource = option
		},
	)

	resourceDropDown.SetCurrentOption(0)
	*selectedResource = "EC2"

	// ==================================================
	// STATUS
	// ==================================================

	statusView := tview.NewTextView()

	statusView.SetText("Status : Ready")
	statusView.SetDynamicColors(true)

	// ==================================================
	// STYLING
	// ==================================================

	dropdowns := []*tview.DropDown{
		profileDropDown,
		regionDropDown,
		resourceDropDown,
	}

	for _, d := range dropdowns {

		d.SetFieldBackgroundColor(tcell.ColorBlack)
		d.SetFieldTextColor(tcell.ColorWhite)
		d.SetLabelColor(tcell.ColorWhite)

		d.SetBackgroundColor(tcell.ColorBlack)
	}

	statusView.SetTextColor(tcell.ColorWhite)
	statusView.SetBackgroundColor(tcell.ColorBlack)

	// ==================================================
	// LEFT COLUMN
	// ==================================================

	leftColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(profileDropDown, 1, 1, true).
		AddItem(regionDropDown, 1, 1, false)

	// ==================================================
	// RIGHT COLUMN
	// ==================================================

	rightColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(resourceDropDown, 1, 1, false).
		AddItem(statusView, 1, 1, false)

	// ==================================================
	// MAIN FLEX
	// ==================================================

	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftColumn, 0, 1, true).
		AddItem(rightColumn, 0, 1, false)

	mainFlex.SetBorder(true)
	mainFlex.SetTitle("A9R")

	return mainFlex,
		profileDropDown,
		regionDropDown,
		resourceDropDown,
		statusView
}