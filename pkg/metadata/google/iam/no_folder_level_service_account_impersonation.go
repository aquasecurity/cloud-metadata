package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoFolderLevelServiceAccountImpersonation = metadata.Metadata{
	ID:          "AVD-GCP-0042",
	Title:       "Users should not be granted service account access at the folder level",
	Description: "Users with service account access at folder level can impersonate any service account. Instead, they should be given access to particular service accounts as required.",
	Impact:      "Privilege escalation, impersonation of any/all services",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/iam/docs/impersonating-service-accounts", 
	},
}

