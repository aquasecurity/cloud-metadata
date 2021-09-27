package rds

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptClusterStorageData = metadata.Metadata{
	ID:          "AVD-AWS-0079",
	Title:       "There is no encryption specified or encryption is disabled on the RDS Cluster.",
	Description: "Encryption should be enabled for an RDS Aurora cluster. 

When enabling encryption by setting the kms_key_id, the storage_encrypted must also be set to true.",
	Impact:      "Data can be read from the RDS cluster if it is compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.html", 
	},
}

