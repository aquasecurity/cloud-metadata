package msk

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableInTransitEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0073",
	Title:       "A MSK cluster allows unencrypted data in transit.",
	Description: "Encryption should be forced for Kafka clusters, including for communication between nodes. This ensure sensitive data is kept private.",
	Impact:      "Intercepted data can be read in transit",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/msk/latest/developerguide/msk-encryption.html", 
	},
}

