package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PgLogDisconnections = metadata.Metadata{
	ID:          "AVD-GCP-0060",
	Title:       "Ensure that logging of disconnections is enabled.",
	Description: "Logging disconnections provides useful diagnostic data such as session length, which can identify performance issues in an application and potential DoS vectors.",
	Impact:      "Insufficient diagnostic data.",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://www.postgresql.org/docs/13/runtime-config-logging.html#GUC-LOG-DISCONNECTIONS", 
	},
}

