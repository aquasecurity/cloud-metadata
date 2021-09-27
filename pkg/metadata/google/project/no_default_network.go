package project

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoDefaultNetwork = metadata.Metadata{
	ID:          "AVD-GCP-0050",
	Title:       "Default network should not be created at project level",
	Description: "The default network which is provided for a project contains multiple insecure firewall rules which allow ingress to the project's infrastructure. Creation of this network should therefore be disabled.",
	Impact:      "Exposure of internal infrastructure/services to public internet",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

