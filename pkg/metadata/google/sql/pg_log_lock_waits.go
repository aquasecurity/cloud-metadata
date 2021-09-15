package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PgLogLockWaits = metadata.Metadata{
	ID:          "AVD-GCP-0062",
	Title:       "Ensure that logging of lock waits is enabled.",
	Description: "Lock waits are often an indication of poor performance and often an indicator of a potential denial of service vulnerability, therefore occurrences should be logged for analysis.",
	Impact:      "Issues leading to denial of service may not be identified.",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://www.postgresql.org/docs/13/runtime-config-logging.html#GUC-LOG-LOCK-WAITS", 
	},
}

