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

	instances := services.GetEC2Instances(
		ctx,
		profile,
		region,
	)

	resource := BuildEC2Resource(instances)

	state.CurrentResource = resource

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
