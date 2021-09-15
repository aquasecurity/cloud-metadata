package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAutoUpgrade = metadata.Metadata{
	ID:          "AVD-GCP-0023",
	Title:       "Kubernetes should have 'Automatic upgrade' enabled",
	Description: "Automatic updates keep nodes updated with the latest cluster master version.",
	Impact:      "Nodes will need the cluster master version manually updating",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

