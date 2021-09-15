package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableBackup = metadata.Metadata{
	ID:          "AVD-GCP-0051",
	Title:       "Enable automated backups to recover from data-loss",
	Description: "Automated backups are not enabled by default. Backups are an easy way to restore data in a corruption or data-loss scenario.",
	Impact:      "No recovery of lost or corrupted data",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/sql/docs/mysql/backup-recovery/backups", 
	},
}

