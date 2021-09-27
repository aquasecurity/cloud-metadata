package elasticache

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableInTransitEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0051",
	Title:       "Elasticache Replication Group uses unencrypted traffic.",
	Description: "Traffic flowing between Elasticache replication nodes should be encrypted to ensure sensitive data is kept private.",
	Impact:      "In transit data in the Replication Group could be read if intercepted",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/in-transit-encryption.html", 
	},
}

