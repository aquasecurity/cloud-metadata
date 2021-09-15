package redshift

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptionCustomerKey = metadata.Metadata{
	ID:          "AVD-AWS-0084",
	Title:       "Redshift clusters should use at rest encryption",
	Description: "Redshift clusters that contain sensitive data or are subject to regulation should be encrypted at rest to prevent data leakage should the infrastructure be compromised.",
	Impact:      "Data may be leaked if infrastructure is compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html", 
	},
}

