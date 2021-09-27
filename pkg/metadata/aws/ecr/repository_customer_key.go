package ecr

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RepositoryCustomerKey = metadata.Metadata{
	ID:          "AVD-AWS-0033",
	Title:       "ECR Repository should use customer managed keys to allow more control",
	Description: "Images in the ECR repository are encrypted by default using AWS managed encryption keys. To increase control of the encryption and control the management of factors like key rotation, use a Customer Managed Key.",
	Impact:      "Using AWS managed keys does not allow for fine grained control",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html", 
	},
}

