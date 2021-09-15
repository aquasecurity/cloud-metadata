package container

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseRbacPermissions = metadata.Metadata{
	ID:          "AVD-AZU-0015",
	Title:       "Ensure RBAC is enabled on AKS clusters",
	Description: "Using Kubernetes role-based access control (RBAC), you can grant users, groups, and service accounts access to only the resources they need.",
	Impact:      "No role based access control is in place for the AKS cluster",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/aks/concepts-identity", 
	},
}

