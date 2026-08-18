package ui

import (
	"context"
	"fmt"
	"strings"

	"github.com/aungshanbo/a9r/aws"
	"github.com/aungshanbo/a9r/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowS3Detail(
	app *tview.Application,
	pages *tview.Pages,
	table *tview.Table,
	profile string,
	detail *models.S3BucketDetail,
) {

	text := tview.NewTextView()

	text.SetDynamicColors(true)
	text.SetBorder(true)
	text.SetTitle(" S3 Bucket Detail ")
	text.SetScrollable(true)

	detail.ObjectCount = -1
	detail.SizeBytes = -1

	render := func() {

		var b strings.Builder

		b.WriteString("\n")

		writeRow := func(label, value string) {
			b.WriteString(
				fmt.Sprintf(
					"  %-20s %s\n",
					label,
					value,
				),
			)
		}

		writeRow("Name", detail.Name)
		writeRow("Region", detail.Region)

		created := "-"

		if !detail.CreationDate.IsZero() {
			created = detail.CreationDate.Format("2006-01-02")
		}

		writeRow("Created", created)

		b.WriteString("\n")

		objectCount := "Loading..."

		if detail.ObjectCount >= 0 {
			objectCount = fmt.Sprintf(
				"%d",
				detail.ObjectCount,
			)
		}

		size := "Loading..."

		if detail.SizeBytes >= 0 {
			size = aws.FormatBytes(detail.SizeBytes)
		}

		writeRow("Objects", objectCount)
		writeRow("Total Size", size)

		b.WriteString("\n")

		writeRow("Versioning", detail.Versioning)
		writeRow("Encryption", detail.Encryption)
		writeRow("Object Lock", detail.ObjectLock)

		b.WriteString("\n")

		writeRow("Public Access", detail.PublicAccess)
		writeRow("Object Ownership", detail.ObjectOwnership)
		writeRow("ACL", detail.ACL)

		b.WriteString("\n")

		writeRow("Policy", detail.Policy)

		writeRow(
			"Lifecycle Rules",
			fmt.Sprintf("%d", detail.LifecycleRules),
		)

		writeRow("Replication", detail.Replication)
		writeRow("Access Logging", detail.AccessLogging)

		b.WriteString("\n")

		writeRow(
			"Tags",
			fmt.Sprintf("%d", len(detail.Tags)),
		)

		b.WriteString("\n")
		b.WriteString("  ESC / q = Close\n")

		text.SetText(b.String())
	}

	render()

	text.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {

			if event.Key() == tcell.KeyEsc ||
				event.Rune() == 'q' {

				pages.RemovePage("s3-detail")
				app.SetFocus(table)

				return nil
			}

			return event
		},
	)

	pages.AddPage(
		"s3-detail",
		text,
		true,
		true,
	)

	app.SetFocus(text)

	// ==========================================
	// BACKGROUND STATISTICS
	// ==========================================

	go func() {

		objectCount, sizeBytes :=
			aws.GetS3BucketStatistics(
				context.Background(),
				profile,
				detail.Region,
				detail.Name,
			)

		app.QueueUpdateDraw(
			func() {

				if !pages.HasPage("s3-detail") {
					return
				}

				detail.ObjectCount = objectCount
				detail.SizeBytes = sizeBytes

				render()
			},
		)
	}()
}
