package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetBuckets(
	ctx context.Context,
	profile string,
	region string,
) []string {

	cfg, err := config.LoadDefaultConfig(
		ctx, config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		return []string{}
	}
	client := s3.NewFromConfig(cfg)

	output, err := client.ListBuckets(
		ctx,
		&s3.ListBucketsInput{},
	)
	if err != nil {
		return []string{}
	}
	var buckets []string

	for _, bucket := range output.Buckets {
		if bucket.Name != nil {
			buckets = append(
				buckets,
				*bucket.Name,
			)
		}
	}
	return buckets
}
