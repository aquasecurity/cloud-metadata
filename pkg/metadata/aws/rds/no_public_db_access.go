package rds

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicDbAccess = metadata.Metadata{
	ID:          "AVD-AWS-0082",
	Title:       "A database resource is marked as publicly accessible.",
	Description: "Database resources should not publicly available. You should limit all access to the minimum that is required for your application to function.",
	Impact:      "The database instance is publicly accessible",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

