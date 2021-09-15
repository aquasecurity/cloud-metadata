package launch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoSensitiveInfo = metadata.Metadata{
	ID:          "AVD-AWS-0068",
	Title:       "Ensure all data stored in the Launch configuration EBS is securely encrypted",
	Description: "When creating Launch Configurations, user data can be used for the initial configuration of the instance. User data must not contain any sensitive data.",
	Impact:      "Sensitive credentials in user data can be leaked",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

