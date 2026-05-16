package aws

import (
	"context"

	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func GetEC2Detail(
	ctx context.Context,
	profile string,
	region string,
	instanceID string,
) *ec2types.Instance {

	cfg, err := awscfg.LoadDefaultConfig(
		ctx,
		awscfg.WithSharedConfigProfile(profile),
		awscfg.WithRegion(region),
	)

	if err != nil {
		return nil
	}

	client := ec2.NewFromConfig(cfg)

	output, err := client.DescribeInstances(
		ctx,
		&ec2.DescribeInstancesInput{
			InstanceIds: []string{
				instanceID,
			},
		},
	)

	if err != nil {
		return nil
	}

	for _, reservation := range output.Reservations {

		for _, instance := range reservation.Instances {

			return &instance
		}
	}

	return nil
}
