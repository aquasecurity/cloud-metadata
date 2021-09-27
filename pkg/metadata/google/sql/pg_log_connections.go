package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PgLogConnections = metadata.Metadata{
	ID:          "AVD-GCP-0059",
	Title:       "Ensure that logging of connections is enabled.",
	Description: "Logging connections provides useful diagnostic data such as session length, which can identify performance issues in an application and potential DoS vectors.",
	Impact:      "Insufficient diagnostic data.",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://www.postgresql.org/docs/13/runtime-config-logging.html#GUC-LOG-CONNECTIONS", 
	},
}

