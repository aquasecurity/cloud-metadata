package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseClusterLabels = metadata.Metadata{
	ID:          "AVD-GCP-0038",
	Title:       "Clusters should be configured with Labels",
	Description: "Labels make it easier to manage assets and differentiate between clusters and environments, allowing the mapping of computational resources to the wider organisational structure.",
	Impact:      "Asset management can be limited/more difficult",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

