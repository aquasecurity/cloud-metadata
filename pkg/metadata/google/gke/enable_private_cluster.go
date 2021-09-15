package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnablePrivateCluster = metadata.Metadata{
	ID:          "AVD-GCP-0027",
	Title:       "Clusters should be set to private",
	Description: "Enabling private nodes on a cluster ensures the nodes are only available internally as they will only be assigned internal addresses.",
	Impact:      "Nodes may be exposed to the public internet",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

