package documentdb

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptionCustomerKey = metadata.Metadata{
	ID:          "AVD-AWS-0022",
	Title:       "DocumentDB encryption should use Customer Managed Keys",
	Description: "Encryption using AWS keys provides protection for your DocumentDB underlying storage. To increase control of the encryption and manage factors like rotation use customer managed keys.",
	Impact:      "Using AWS managed keys does not allow for fine grained control",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

