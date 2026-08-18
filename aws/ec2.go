package aws

import (
	"context"

	"github.com/aungshanbo/a9r/models"
	"github.com/aws/aws-sdk-go-v2/config"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func GetEC2Instances(
	ctx context.Context,
	profile string,
	region string,
) []models.Ec2instance {

	var instances []models.Ec2instance

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)

	if err != nil {
		return instances
	}

	client := ec2.NewFromConfig(cfg)

	result, err := client.DescribeInstances(
		ctx,
		&ec2.DescribeInstancesInput{},
	)

	if err != nil {
		return instances
	}

	for _, r := range result.Reservations {

		for _, inst := range r.Instances {

			id := "-"
			state := "-"
			itype := "-"
			name := "-"
			az := "-"
			privateIP := "-"
			publicIP := "-"

			if inst.InstanceId != nil {
				id = *inst.InstanceId
			}

			if inst.State != nil {
				state = string(inst.State.Name)
			}

			if inst.InstanceType != "" {
				itype = string(inst.InstanceType)
			}

			for _, tag := range inst.Tags {

				if tag.Key != nil &&
					*tag.Key == "Name" {

					name = *tag.Value
				}
			}

			if inst.PrivateIpAddress != nil {
				privateIP = *inst.PrivateIpAddress
			}

			if inst.PublicIpAddress != nil {
				publicIP = *inst.PublicIpAddress
			}

			if inst.Placement != nil &&
				inst.Placement.AvailabilityZone != nil {

				az = *inst.Placement.AvailabilityZone
			}

			instances = append(
				instances,
				models.Ec2instance{
					ID:        id,
					State:     state,
					Type:      itype,
					Name:      name,
					AZ:        az,
					PublicIP:  publicIP,
					PrivateIP: privateIP,
				},
			)
		}
	}

	return instances
}

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
