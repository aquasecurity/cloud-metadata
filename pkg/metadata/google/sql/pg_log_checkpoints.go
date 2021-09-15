package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PgLogCheckpoints = metadata.Metadata{
	ID:          "AVD-GCP-0058",
	Title:       "Ensure that logging of checkpoints is enabled.",
	Description: "Logging checkpoints provides useful diagnostic data, which can identify performance issues in an application and potential DoS vectors.",
	Impact:      "Insufficient diagnostic data.",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://www.postgresql.org/docs/13/runtime-config-logging.html#GUC-LOG-CHECKPOINTS", 
	},
}

