package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoSensitiveInfo = metadata.Metadata{
	ID:          "AVD-CSK-0001",
	Title:       "No sensitive data stored in user_data",
	Description: "When creating instances, user data can be used during the initial configuration. User data must not contain sensitive information",
	Impact:      "Sensitive credentials in the user data can be leaked",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

