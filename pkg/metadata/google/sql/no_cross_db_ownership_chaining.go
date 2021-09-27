package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoCrossDbOwnershipChaining = metadata.Metadata{
	ID:          "AVD-GCP-0056",
	Title:       "Cross-database ownership chaining should be disabled",
	Description: "Cross-database ownership chaining, also known as cross-database chaining, is a security feature of SQL Server that allows users of databases access to other databases besides the one they are currently using.",
	Impact:      "Unintended access to sensitive data",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/cross-db-ownership-chaining-server-configuration-option?view=sql-server-ver15", 
	},
}

