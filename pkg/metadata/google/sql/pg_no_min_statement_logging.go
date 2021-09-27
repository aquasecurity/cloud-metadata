package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PgNoMinStatementLogging = metadata.Metadata{
	ID:          "AVD-GCP-0063",
	Title:       "Ensure that logging of long statements is disabled.",
	Description: "Logging of statements which could contain sensitive data is not advised, therefore this setting should preclude all statements from being logged.",
	Impact:      "Sensitive data could be exposed in the database logs.",
	Severity:    "LOW",
	Links:       []string {
		"https://www.postgresql.org/docs/13/runtime-config-logging.html#GUC-LOG-MIN-DURATION-STATEMENT", 
	},
}

