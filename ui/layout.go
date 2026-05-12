package ui

import "github.com/rivo/tview"

func NewLayout(
	leftPanel tview.Primitive,
	table *tview.Table,
	status *tview.TextView,
) *tview.Flex {

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(leftPanel, 7, 1, true).
		AddItem(table, 0, 3, false).
		AddItem(status, 3, 1, false)
}
