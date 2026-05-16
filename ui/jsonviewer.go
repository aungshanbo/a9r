package ui

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func highlightSearchText(
	content string,
	search string,
	currentMatch int,
) string {

	if search == "" {
		return content
	}

	lowerContent := strings.ToLower(content)
	lowerSearch := strings.ToLower(search)

	result := ""

	i := 0
	matchIndex := 0

	for i < len(content) {

		index := strings.Index(
			lowerContent[i:],
			lowerSearch,
		)

		if index == -1 {

			result += content[i:]

			break
		}

		index += i

		result += content[i:index]

		match :=
			content[index : index+len(search)]

		// current focused match
		if matchIndex == currentMatch {

			result +=
				"[white:red]" +
					match +
					"[-:-]"

		} else {

			result +=
				"[black:yellow]" +
					match +
					"[-:-]"
		}

		matchIndex++

		i = index + len(search)
	}

	return result
}

func scrollToMatch(
	text *tview.TextView,
	line int,
) {

	matchLine := line

	if matchLine > 10 {
		matchLine -= 10
	}

	text.ScrollTo(
		matchLine,
		0,
	)
}

func updateSearchTitle(
	text *tview.TextView,
	title string,
	current int,
	total int,
	search string,
) {

	text.SetTitle(
		fmt.Sprintf(
			"%s | %d/%d : %s",
			title,
			current,
			total,
			search,
		),
	)
}

func ShowJSONViewer(
	app *tview.Application,
	pages *tview.Pages,
	table *tview.Table,
	title string,
	data interface{},
) {

	searchMode := false
	searchText := ""

	searchResults := []int{}
	currentMatch := 0

	jsonData, err := json.MarshalIndent(
		data,
		"",
		"  ",
	)

	if err != nil {
		return
	}

	jsonString := string(jsonData)

	lines := strings.Split(
		jsonString,
		"\n",
	)

	text := tview.NewTextView().
		SetText(jsonString)

	text.SetBorder(true)
	text.SetTitle(title)
	text.SetScrollable(true)
	text.SetDynamicColors(true)

	modal := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(text, 0, 1, true)

	text.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// ==========================================
		// SEARCH MODE
		// ==========================================
		if searchMode {

			switch {

			// execute search
			case event.Key() == tcell.KeyEnter:

				searchResults = []int{}
				currentMatch = 0

				for i, line := range lines {

					if strings.Contains(
						strings.ToLower(line),
						strings.ToLower(searchText),
					) {

						searchResults = append(
							searchResults,
							i,
						)
					}
				}

				if len(searchResults) > 0 {

					text.SetText(
						highlightSearchText(
							jsonString,
							searchText,
							currentMatch,
						),
					)

					scrollToMatch(
						text,
						searchResults[currentMatch],
					)

					updateSearchTitle(
						text,
						title,
						currentMatch+1,
						len(searchResults),
						searchText,
					)

				} else {

					text.SetTitle(
						fmt.Sprintf(
							"%s | No Match : %s",
							title,
							searchText,
						),
					)
				}

				searchMode = false

				return nil

			// cancel search mode
			case event.Key() == tcell.KeyEsc:

				searchMode = false
				searchText = ""

				text.SetText(jsonString)

				text.SetTitle(title)

				return nil

			// backspace
			case event.Key() == tcell.KeyBackspace,
				event.Key() == tcell.KeyBackspace2:

				if len(searchText) > 0 {

					searchText =
						searchText[:len(searchText)-1]
				}

			// typing
			case event.Rune() != 0:

				searchText += string(event.Rune())
			}

			text.SetTitle(
				title +
					" | /" +
					searchText,
			)

			return nil
		}

		// ==========================================
		// NORMAL MODE
		// ==========================================
		switch {

		// search mode
		case event.Rune() == '/':

			searchMode = true
			searchText = ""

			text.SetTitle(
				title +
					" | /",
			)

			return nil

		// next search result
		case event.Key() == tcell.KeyEnter:

			if len(searchResults) == 0 {
				return nil
			}

			text.SetText(
				highlightSearchText(
					jsonString,
					searchText,
					currentMatch,
				),
			)

			scrollToMatch(
				text,
				searchResults[currentMatch],
			)

			updateSearchTitle(
				text,
				title,
				currentMatch+1,
				len(searchResults),
				searchText,
			)

			currentMatch++

			if currentMatch >= len(searchResults) {
				currentMatch = 0
			}

			return nil

		// close modal
		case event.Key() == tcell.KeyEsc:

			pages.RemovePage("json-viewer")

			app.SetFocus(table)

			return nil

		// vim down
		case event.Rune() == 'j':

			row, col := text.GetScrollOffset()

			text.ScrollTo(row+1, col)

			return nil

		// vim up
		case event.Rune() == 'k':

			row, col := text.GetScrollOffset()

			if row > 0 {
				text.ScrollTo(row-1, col)
			}

			return nil

		// arrow down
		case event.Key() == tcell.KeyDown:

			row, col := text.GetScrollOffset()

			text.ScrollTo(row+1, col)

			return nil

		// arrow up
		case event.Key() == tcell.KeyUp:

			row, col := text.GetScrollOffset()

			if row > 0 {
				text.ScrollTo(row-1, col)
			}

			return nil
		}

		return event
	})

	pages.AddPage(
		"json-viewer",
		Center(140, 40, modal),
		true,
		true,
	)

	app.SetFocus(text)
}
