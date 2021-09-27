package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIp = metadata.Metadata{
	ID:          "AVD-GCP-0015",
	Title:       "Instances should not have public IP addresses",
	Description: "Instances should not be publicly exposed to the internet",
	Impact:      "Direct exposure of an instance to the public internet",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

