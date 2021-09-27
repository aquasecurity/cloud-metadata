package dynamodb

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var TableCustomerKey = metadata.Metadata{
	ID:          "AVD-AWS-0025",
	Title:       "DynamoDB tables should use at rest encryption with a Customer Managed Key",
	Description: "DynamoDB tables are encrypted by default using AWS managed encryption keys. To increase control of the encryption and control the management of factors like key rotation, use a Customer Managed Key.",
	Impact:      "Using AWS managed keys does not allow for fine grained control",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/EncryptionAtRest.html", 
	},
}

