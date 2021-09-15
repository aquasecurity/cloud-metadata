package athena

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAtRestEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0006",
	Title:       "Athena databases and workgroup configurations are created unencrypted at rest by default, they should be encrypted",
	Description: "Athena databases and workspace result sets should be encrypted at rests. These databases and query sets are generally derived from data in S3 buckets and should have the same level of at rest protection.",
	Impact:      "Data can be read if the Athena Database is compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/athena/latest/ug/encryption.html", 
	},
}

