package services

import (
	"context"

	"github.com/aungshanbo/a9r/aws"
)

func GetS3Buckets(
	ctx context.Context,
	profile string,
	region string,
) []string {
	return aws.Getbuckets(
		ctx,
		profile,
		region,
	)
}
