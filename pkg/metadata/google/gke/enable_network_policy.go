package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableNetworkPolicy = metadata.Metadata{
	ID:          "AVD-GCP-0026",
	Title:       "Network Policy should be enabled on GKE clusters",
	Description: "Enabling a network policy allows the segregation of network traffic by namespace",
	Impact:      "Unrestricted inter-cluster communication",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

