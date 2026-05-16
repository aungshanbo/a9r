package services

import (
	"context"

	"github.com/aungshanbo/a9r/aws"
	"github.com/aungshanbo/a9r/models"
)

func GetS3Buckets(
	ctx context.Context,
	profile string,
	region string,
) []string {
	return aws.GetBuckets(
		ctx,
		profile,
		region,
	)
}

func BuildS3Resource(
	buckets []string,
) *models.Resource {

	headers := []models.TableColumn{
		{
			Title:     "Bucket Name",
			Expansion: 1,
		},
	}

	var rows [][]string
	for _, bucket := range buckets {
		rows = append(rows, []string{
			bucket,
		})
	}
	return &models.Resource{
		Headers: headers,
		Rows:    rows,
	}
}
