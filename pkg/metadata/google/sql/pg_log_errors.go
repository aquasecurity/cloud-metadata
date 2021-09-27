package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var PgLogErrors = metadata.Metadata{
	ID:          "AVD-GCP-0061",
	Title:       "Ensure that Postgres errors are logged",
	Description: "Setting the minimum log severity too high will cause errors not to be logged",
	Impact:      "Loss of error logging",
	Severity:    "LOW",
	Links:       []string {
		"https://postgresqlco.nf/doc/en/param/log_min_messages/", "https://www.postgresql.org/docs/13/runtime-config-logging.html#GUC-LOG-MIN-MESSAGES", 
	},
}

