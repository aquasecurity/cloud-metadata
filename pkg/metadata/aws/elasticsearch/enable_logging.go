package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableLogging = metadata.Metadata{
	ID:          "AVD-AWS-0044",
	Title:       "AWS ES Domain should have logging enabled",
	Description: "AWS ES domain should have logging enabled by default.",
	Impact:      "Logging provides vital information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

