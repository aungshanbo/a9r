package aws

import (
	"context"
	"fmt"

	"github.com/aungshanbo/a9r/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetBuckets(
	ctx context.Context,
	profile string,
	region string,
) []models.S3Bucket {

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		return []models.S3Bucket{}
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListBuckets(
		ctx,
		&s3.ListBucketsInput{},
	)
	if err != nil {
		return []models.S3Bucket{}
	}

	buckets := make([]models.S3Bucket, 0, len(output.Buckets))

	for _, bucket := range output.Buckets {
		if bucket.Name == nil {
			continue
		}

		item := models.S3Bucket{
			Name: *bucket.Name,
		}

		if bucket.CreationDate != nil {
			item.CreationDate = *bucket.CreationDate
		}

		buckets = append(buckets, item)
	}

	return buckets
}

func GetS3BucketDetail(
	ctx context.Context,
	profile string,
	region string,
	bucketName string,
) *models.S3BucketDetail {

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(cfg)

	detail := &models.S3BucketDetail{
		Name:            bucketName,
		Region:          region,
		Versioning:      "Unknown",
		Encryption:      "Unknown",
		ObjectLock:      "Unknown",
		PublicAccess:    "Unknown",
		ObjectOwnership: "Unknown",
		ACL:             "Unknown",
		Policy:          "Unknown",
		Replication:     "Unknown",
		AccessLogging:   "Unknown",
		Tags:            map[string]string{},
	}

	// Versioning
	versioning, err := client.GetBucketVersioning(
		ctx,
		&s3.GetBucketVersioningInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil {
		switch versioning.Status {
		case "Enabled":
			detail.Versioning = "Enabled"
		case "Suspended":
			detail.Versioning = "Suspended"
		default:
			detail.Versioning = "Disabled"
		}
	}

	// Encryption
	encryption, err := client.GetBucketEncryption(
		ctx,
		&s3.GetBucketEncryptionInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil &&
		encryption.ServerSideEncryptionConfiguration != nil {

		rules := encryption.ServerSideEncryptionConfiguration.Rules

		if len(rules) > 0 &&
			rules[0].ApplyServerSideEncryptionByDefault != nil {

			defaultEncryption :=
				rules[0].ApplyServerSideEncryptionByDefault

			detail.Encryption =
				string(defaultEncryption.SSEAlgorithm)

			if defaultEncryption.KMSMasterKeyID != nil {
				detail.Encryption = "SSE-KMS"
			}
		}
	}

	if detail.Encryption == "Unknown" {
		detail.Encryption = "SSE-S3"
	}

	// Object Lock
	objectLock, err := client.GetObjectLockConfiguration(
		ctx,
		&s3.GetObjectLockConfigurationInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil &&
		objectLock.ObjectLockConfiguration != nil {

		if objectLock.ObjectLockConfiguration.ObjectLockEnabled == "Enabled" {
			detail.ObjectLock = "Enabled"
		} else {
			detail.ObjectLock = "Disabled"
		}
	} else {
		detail.ObjectLock = "Disabled"
	}

	// Public Access Block
	publicAccess, err := client.GetPublicAccessBlock(
		ctx,
		&s3.GetPublicAccessBlockInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil &&
		publicAccess.PublicAccessBlockConfiguration != nil {

		block := publicAccess.PublicAccessBlockConfiguration

		if aws.ToBool(block.BlockPublicAcls) &&
			aws.ToBool(block.IgnorePublicAcls) &&
			aws.ToBool(block.BlockPublicPolicy) &&
			aws.ToBool(block.RestrictPublicBuckets) {

			detail.PublicAccess = "Blocked"
		} else {
			detail.PublicAccess = "Partial"
		}
	}

	// Object Ownership
	ownership, err := client.GetBucketOwnershipControls(
		ctx,
		&s3.GetBucketOwnershipControlsInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil &&
		ownership.OwnershipControls != nil {

		rules := ownership.OwnershipControls.Rules

		if len(rules) > 0 {
			detail.ObjectOwnership =
				string(rules[0].ObjectOwnership)
		}
	}

	// ACL
	acl, err := client.GetBucketAcl(
		ctx,
		&s3.GetBucketAclInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil && acl.Owner != nil {
		detail.ACL = "private"
	}

	// Policy
	policy, err := client.GetBucketPolicyStatus(
		ctx,
		&s3.GetBucketPolicyStatusInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil && policy.PolicyStatus != nil {

		detail.Policy = "Configured"

		if aws.ToBool(policy.PolicyStatus.IsPublic) {
			detail.Policy = "PUBLIC"
		}
	} else {
		detail.Policy = "Not configured"
	}

	// Lifecycle
	lifecycle, err := client.GetBucketLifecycleConfiguration(
		ctx,
		&s3.GetBucketLifecycleConfigurationInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil {
		detail.LifecycleRules = len(lifecycle.Rules)
	}

	// Replication
	replication, err := client.GetBucketReplication(
		ctx,
		&s3.GetBucketReplicationInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil &&
		replication.ReplicationConfiguration != nil {

		detail.Replication = "Enabled"
	} else {
		detail.Replication = "Disabled"
	}

	// Access Logging
	logging, err := client.GetBucketLogging(
		ctx,
		&s3.GetBucketLoggingInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil && logging.LoggingEnabled != nil {
		detail.AccessLogging = "Enabled"
	} else {
		detail.AccessLogging = "Disabled"
	}

	// Tags
	tagging, err := client.GetBucketTagging(
		ctx,
		&s3.GetBucketTaggingInput{
			Bucket: aws.String(bucketName),
		},
	)

	if err == nil {

		for _, tag := range tagging.TagSet {

			if tag.Key == nil {
				continue
			}

			value := ""

			if tag.Value != nil {
				value = *tag.Value
			}

			detail.Tags[*tag.Key] = value
		}
	}

	return detail
}

func GetS3BucketStatistics(
	ctx context.Context,
	profile string,
	region string,
	bucketName string,
) (int64, int64) {

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		return 0, 0
	}

	client := s3.NewFromConfig(cfg)

	var objectCount int64
	var sizeBytes int64
	var token *string

	for {

		output, err := client.ListObjectsV2(
			ctx,
			&s3.ListObjectsV2Input{
				Bucket:            aws.String(bucketName),
				ContinuationToken: token,
			},
		)

		if err != nil {
			return objectCount, sizeBytes
		}

		for _, object := range output.Contents {

			objectCount++

			if object.Size != nil {
				sizeBytes += *object.Size
			}
		}

		if !aws.ToBool(output.IsTruncated) {
			break
		}

		token = output.NextContinuationToken

		if token == nil {
			break
		}
	}

	return objectCount, sizeBytes
}

func FormatBytes(bytes int64) string {

	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}

	value := float64(bytes)

	units := []string{
		"KB",
		"MB",
		"GB",
		"TB",
		"PB",
	}

	for _, unit := range units {

		value /= 1024

		if value < 1024 {
			return fmt.Sprintf("%.2f %s", value, unit)
		}
	}

	return fmt.Sprintf("%.2f EB", value)
}
