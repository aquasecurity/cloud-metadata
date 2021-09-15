package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PostgresConfigurationLogConnections = metadata.Metadata{
	ID:          "AVD-AZU-0022",
	Title:       "Ensure server parameter 'log_connections' is set to 'ON' for PostgreSQL Database Server",
	Description: "Postgresql can generate logs for successful connections to improve visibility for audit and configuration issue resolution.",
	Impact:      "No visibility of successful connections",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/postgresql/concepts-server-logs#configure-logging", 
	},
}

