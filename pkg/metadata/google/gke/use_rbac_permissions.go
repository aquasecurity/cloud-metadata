package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseRbacPermissions = metadata.Metadata{
	ID:          "AVD-GCP-0039",
	Title:       "Legacy ABAC permissions are enabled.",
	Description: "You should disable Attribute-Based Access Control (ABAC), and instead use Role-Based Access Control (RBAC) in GKE.

RBAC has significant security advantages and is now stable in Kubernetes, so it’s time to disable ABAC.",
	Impact:      "ABAC permissions are less secure than RBAC permissions",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#leave_abac_disabled_default_for_110", 
	},
}

