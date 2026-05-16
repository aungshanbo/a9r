package services

import (
	"context"

	"github.com/aungshanbo/a9r/aws"
	"github.com/aungshanbo/a9r/models"
)

func GetEC2Instances(
	ctx context.Context,
	profile string,
	region string,
) []models.Ec2instance {

	instances := aws.GetEC2Instances(
		ctx,
		profile,
		region,
	)

	return instances
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
