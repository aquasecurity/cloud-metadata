package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptReplicationGroup = metadata.Metadata{
	ID:          "AVD-AWS-0045",
	Title:       "Unencrypted Elasticache Replication Group.",
	Description: "You should ensure your Elasticache data is encrypted at rest to help prevent sensitive information from being read by unauthorised users.",
	Impact:      "Data in the replication group could be readable if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/at-rest-encryption.html", 
	},
}

