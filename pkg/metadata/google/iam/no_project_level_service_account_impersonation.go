package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoProjectLevelServiceAccountImpersonation = metadata.Metadata{
	ID:          "AVD-GCP-0047",
	Title:       "Users should not be granted service account access at the project level",
	Description: "Users with service account access at project level can impersonate any service account. Instead, they should be given access to particular service accounts as required.",
	Impact:      "Privilege escalation, impersonation of any/all services",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/iam/docs/impersonating-service-accounts", 
	},
}

