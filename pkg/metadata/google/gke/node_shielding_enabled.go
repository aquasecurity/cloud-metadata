package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NodeShieldingEnabled = metadata.Metadata{
	ID:          "AVD-GCP-0037",
	Title:       "Shielded GKE nodes not enabled.",
	Description: "CIS GKE Benchmark Recommendation: 6.5.5. Ensure Shielded GKE Nodes are Enabled

Shielded GKE Nodes provide strong, verifiable node identity and integrity to increase the security of GKE nodes and should be enabled on all GKE clusters.",
	Impact:      "Node identity and integrity can't be verified without shielded GKE nodes",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#shielded_nodes", 
	},
}

