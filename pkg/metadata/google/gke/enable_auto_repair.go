package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAutoRepair = metadata.Metadata{
	ID:          "AVD-GCP-0022",
	Title:       "Kubernetes should have 'Automatic repair' enabled",
	Description: "Automatic repair will monitor nodes and attempt repair when a node fails multiple subsequent health checks",
	Impact:      "Failing nodes will require manual repair.",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

