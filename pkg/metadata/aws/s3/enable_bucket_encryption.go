package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableBucketEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0088",
	Title:       "Unencrypted S3 bucket.",
	Description: "S3 Buckets should be encrypted with customer managed KMS keys and not default AWS managed keys, in order to allow granular control over access to specific buckets.",
	Impact:      "The bucket objects could be read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucket-encryption.html", 
	},
}

