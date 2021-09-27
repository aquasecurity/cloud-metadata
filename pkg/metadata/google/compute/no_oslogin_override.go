package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoOsloginOverride = metadata.Metadata{
	ID:          "AVD-GCP-0009",
	Title:       "Instances should not override the project setting for OS Login",
	Description: "OS Login automatically revokes the relevant SSH keys when an IAM user has their access revoked.",
	Impact:      "Access via SSH key cannot be revoked automatically when an IAM user is removed.",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

