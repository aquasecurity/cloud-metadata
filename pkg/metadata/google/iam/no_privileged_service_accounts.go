package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPrivilegedServiceAccounts = metadata.Metadata{
	ID:          "AVD-GCP-0045",
	Title:       "Service accounts should not have roles assigned with excessive privileges",
	Description: "Service accounts should have a minimal set of permissions assigned in order to do their job. They should never have excessive access as if compromised, an attacker can escalate privileges and take over the entire account.",
	Impact:      "Cloud account takeover if a resource using a service account is compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/iam/docs/understanding-roles", 
	},
}

