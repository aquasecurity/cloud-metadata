package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-AZU-0018",
	Title:       "Ensure databases are not publicly accessible",
	Description: "Database resources should not publicly available. You should limit all access to the minimum that is required for your application to function.",
	Impact:      "Publicly accessible database could lead to compromised data",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

