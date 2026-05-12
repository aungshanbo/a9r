package ui

import "strings"

func FilterRows(
	rows [][]string,
	filter string,
) [][]string {
	if filter == "" {
		return rows
	}
	filter = strings.ToLower(filter)
	var filtered [][]string
	for _, row := range rows {
		for _, col := range row {
			if strings.Contains(
				strings.ToLower(col),
				filter,
			) {
				filtered = append(filtered, row)
				break
			}
		}
	}
	return filtered
}
