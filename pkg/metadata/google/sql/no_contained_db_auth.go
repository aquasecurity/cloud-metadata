package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoContainedDbAuth = metadata.Metadata{
	ID:          "AVD-GCP-0055",
	Title:       "Contained database authentication should be disabled",
	Description: "Users with ALTER permissions on users can grant access to a contained database without the knowledge of an administrator",
	Impact:      "Access can be granted without knowledge of the database administrator",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/contained-database-authentication-server-configuration-option?view=sql-server-ver15", 
	},
}

