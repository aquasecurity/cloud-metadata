package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PostgresConfigurationLogConnectionThrottling = metadata.Metadata{
	ID:          "AVD-AZU-0021",
	Title:       "Ensure server parameter 'connection_throttling' is set to 'ON' for PostgreSQL Database Server",
	Description: "Postgresql can generate logs for connection throttling to improve visibility for audit and configuration issue resolution.",
	Impact:      "No log information to help diagnosing connection contention issues",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/postgresql/concepts-server-logs#configure-logging", 
	},
}

