package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NodePoolUsesCos = metadata.Metadata{
	ID:          "AVD-GCP-0036",
	Title:       "Ensure Container-Optimized OS (cos) is used for Kubernetes Engine Clusters Node image",
	Description: "GKE supports several OS image types but COS is the recommended OS image to use on cluster nodes for enhanced security",
	Impact:      "COS is the recommended OS image to use on cluster nodes",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

