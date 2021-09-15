package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var ProjectLevelOslogin = metadata.Metadata{
	ID:          "AVD-GCP-0017",
	Title:       "OS Login should be enabled at project level",
	Description: "OS Login automatically revokes the relevant SSH keys when an IAM user has their access revoked.",
	Impact:      "Access via SSH key cannot be revoked automatically when an IAM user is removed.",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

