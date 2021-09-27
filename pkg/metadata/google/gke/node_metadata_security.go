package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NodeMetadataSecurity = metadata.Metadata{
	ID:          "AVD-GCP-0035",
	Title:       "Node metadata value disables metadata concealment.",
	Description: "If the workload_metadata_config block within node_config is included, the node_metadata attribute should be configured securely.

The attribute should be set to SECURE to use metadata concealment, or GKE_METADATA_SERVER if workload identity is enabled. This ensures that the VM metadata is not unnecessarily exposed to pods.",
	Impact:      "Metadata that isn't concealed potentially risks leakage of sensitive data",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/protecting-cluster-metadata#create-concealed", 
	},
}

