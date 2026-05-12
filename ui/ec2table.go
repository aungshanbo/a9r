package ui

import (
	"github.com/aungshanbo/a9r/models"
	"github.com/aungshanbo/a9r/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DrawEC2Table(
	table *tview.Table,
	instances []models.Ec2instance,
) {

	table.Clear()

	var EC2Columns = []models.TableColumn{
		{Title: "Name", Expansion: 2},
		{Title: "ID", Expansion: 3},
		{Title: "State", Expansion: 1},
		{Title: "Type", Expansion: 1},
		{Title: "AZ", Expansion: 2},
		{Title: "Private IP", Expansion: 2},
		{Title: "Public IP", Expansion: 2},
	}

	for i, h := range EC2Columns {

		table.SetCell(
			0,
			i,
			tview.NewTableCell(h.Title).
				SetAttributes(tcell.AttrBold).
				SetExpansion(h.Expansion),
		)
	}

	for i, inst := range instances {

		color := utils.Getstatecolor(inst.State)

		table.SetCell(i+1, 0,
			tview.NewTableCell(inst.Name))

		table.SetCell(i+1, 1,
			tview.NewTableCell(inst.ID))

		table.SetCell(i+1, 2,
			tview.NewTableCell(inst.State).
				SetTextColor(color))

		table.SetCell(i+1, 3,
			tview.NewTableCell(inst.Type))

		table.SetCell(i+1, 4,
			tview.NewTableCell(inst.AZ))

		table.SetCell(i+1, 5,
			tview.NewTableCell(inst.PrivateIP))

		table.SetCell(i+1, 6,
			tview.NewTableCell(inst.PublicIP))
	}
}

func BuildEC2Resource(
	instances []models.Ec2instance,
) *models.Resource {

	rows := [][]string{}

	for _, inst := range instances {

		rows = append(rows, []string{
			inst.ID,
			inst.Name,
			inst.State,
			inst.Type,
			inst.AZ,
			inst.PrivateIP,
			inst.PublicIP,
		})
	}

	return &models.Resource{
		Name: "EC2",
		Headers: []models.TableColumn{
			{Title: "ID", Expansion: 2},
			{Title: "Name", Expansion: 3},
			{Title: "State", Expansion: 1},
			{Title: "Type", Expansion: 1},
			{Title: "AZ", Expansion: 2},
			{Title: "Private IP", Expansion: 2},
			{Title: "Public IP", Expansion: 2},
		},
		Rows: rows,
	}
}
