package ui

import (
	"context"
	"sync"

	"github.com/aungshanbo/a9r/models"
	"github.com/aungshanbo/a9r/services"
	"github.com/rivo/tview"
)

var mu sync.Mutex
var cancelFunc context.CancelFunc

func RefreshTable(
	app *tview.Application,
	table *tview.Table,
	statusView *tview.TextView,
	state *models.AppState,
	profile string,
	region string,
	resourceType string,
) {

	mu.Lock()

	if cancelFunc != nil {
		cancelFunc()
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancelFunc = cancel

	mu.Unlock()

	app.QueueUpdateDraw(func() {
		statusView.SetText("Status : Loading...")
	})

	var resource *models.Resource

	switch resourceType {
	case "EC2":
		instanes := services.GetEC2Instances(
			ctx,
			profile,
			region)
		resource = BuildEC2Resource(instanes)

	case "S3":
		buckets := services.GetS3Buckets(
			ctx,
			profile,
			region)
		resource = BuildS3Resource(buckets)
	}

	state.CurrentResource = resource
	state.ResourceType = resourceType

	state.CurrentResource.Rows = FilterRows(
		state.CurrentResource.Rows,
		state.Filter,
	)

	select {

	case <-ctx.Done():
		return

	default:
	}

	app.QueueUpdateDraw(func() {

		DrawTable(
			table,
			state.CurrentResource.Headers,
			state.CurrentResource.Rows,
		)

		statusView.SetText("Status : Done ✓")
	})
}
