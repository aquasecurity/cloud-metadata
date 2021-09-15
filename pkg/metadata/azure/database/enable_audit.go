package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAudit = metadata.Metadata{
	ID:          "AVD-AZU-0016",
	Title:       "Auditing should be enabled on Azure SQL Databases",
	Description: "Auditing helps you maintain regulatory compliance, understand database activity, and gain insight into discrepancies and anomalies that could indicate business concerns or suspected security violations.",
	Impact:      "Auditing provides valuable information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/azure-sql/database/auditing-overview", 
	},
}

