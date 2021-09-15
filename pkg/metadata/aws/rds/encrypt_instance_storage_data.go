package rds

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptInstanceStorageData = metadata.Metadata{
	ID:          "AVD-AWS-0080",
	Title:       "RDS encryption has not been enabled at a DB Instance level.",
	Description: "Encryption should be enabled for an RDS Database instances. 

When enabling encryption by setting the kms_key_id.",
	Impact:      "Data can be read from RDS instances if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.html", 
	},
}

