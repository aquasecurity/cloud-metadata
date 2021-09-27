package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoOrgLevelServiceAccountImpersonation = metadata.Metadata{
	ID:          "AVD-GCP-0044",
	Title:       "Users should not be granted service account access at the organization level",
	Description: "Users with service account access at organization level can impersonate any service account. Instead, they should be given access to particular service accounts as required.",
	Impact:      "Privilege escalation, impersonation of any/all services",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/iam/docs/impersonating-service-accounts", 
	},
}

