package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RetentionPeriodSet = metadata.Metadata{
	ID:          "AVD-AZU-0023",
	Title:       "Database auditing rentention period should be longer than 90 days",
	Description: "When Auditing is configured for a SQL database, if the retention period is not set, the retention will be unlimited.

If the retention period is to be explicitly set, it should be set for no less than 90 days.",
	Impact:      "Short logging retention could result in missing valuable historical information",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/azure-sql/database/auditing-overview", 
	},
}

