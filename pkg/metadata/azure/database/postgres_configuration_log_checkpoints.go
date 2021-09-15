package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PostgresConfigurationLogCheckpoints = metadata.Metadata{
	ID:          "AVD-AZU-0020",
	Title:       "Ensure server parameter 'log_checkpoints' is set to 'ON' for PostgreSQL Database Server",
	Description: "Postgresql can generate logs for checkpoints to improve visibility for audit and configuration issue resolution.",
	Impact:      "No error and query logs generated on checkpoint",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/postgresql/concepts-server-logs#configure-logging", 
	},
}

