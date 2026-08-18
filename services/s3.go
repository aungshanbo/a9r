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
) []models.S3Bucket {

	return aws.GetBuckets(
		ctx,
		profile,
		region,
	)
}

func BuildS3Resource(
	buckets []models.S3Bucket,
) *models.Resource {

	headers := []models.TableColumn{
		{
			Title:     "Bucket Name",
			Expansion: 3,
		},
		{
			Title:     "Created",
			Expansion: 2,
		},
	}

	rows := make([][]string, 0, len(buckets))

	for _, bucket := range buckets {

		created := "-"

		if !bucket.CreationDate.IsZero() {
			created = bucket.CreationDate.Format("2006-01-02")
		}

		rows = append(
			rows,
			[]string{
				bucket.Name,
				created,
			},
		)
	}

	return &models.Resource{
		Name:    "S3",
		Headers: headers,
		Rows:    rows,
	}
}
