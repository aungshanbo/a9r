package utils

import "github.com/gdamore/tcell/v2"

func Getstatecolor(state string) tcell.Color {
	switch state {
	case "running":
		return tcell.ColorGreen
	case "stopped":
		return tcell.ColorRed
	case "pending":
		return tcell.ColorYellow
	case "stopping":
		return tcell.ColorOrange
	default:
		return tcell.ColorWhite
	}
}
