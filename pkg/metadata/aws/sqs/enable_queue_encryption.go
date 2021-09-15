package sqs

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableQueueEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0096",
	Title:       "Unencrypted SQS queue.",
	Description: "Queues should be encrypted with customer managed KMS keys and not default AWS managed keys, in order to allow granular control over access to specific queues.",
	Impact:      "The SQS queue messages could be read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html", 
	},
}

