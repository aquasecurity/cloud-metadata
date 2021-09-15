package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableMasterNetworks = metadata.Metadata{
	ID:          "AVD-GCP-0025",
	Title:       "Master authorized networks should be configured on GKE clusters",
	Description: "Enabling authorized networks means you can restrict master access to a fixed set of CIDR ranges",
	Impact:      "Unrestricted network access to the master",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

