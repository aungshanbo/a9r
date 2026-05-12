package utils

import "fmt"

func CenterText(text string, width int) string {

	padding := (width - len(text)) / 2

	if padding < 0 {
		padding = 0
	}

	return fmt.Sprintf("%*s%s", padding, "", text)
}
