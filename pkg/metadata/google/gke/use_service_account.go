package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseServiceAccount = metadata.Metadata{
	ID:          "AVD-GCP-0040",
	Title:       "Checks for service account defined for GKE nodes",
	Description: "You should create and use a minimally privileged service account to run your GKE cluster instead of using the Compute Engine default service account.",
	Impact:      "Service accounts with wide permissions can increase the risk of compromise",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#use_least_privilege_sa", 
	},
}

