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
