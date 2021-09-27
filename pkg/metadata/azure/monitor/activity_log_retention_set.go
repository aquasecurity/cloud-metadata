package monitor

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var ActivityLogRetentionSet = metadata.Metadata{
	ID:          "AVD-AZU-0033",
	Title:       "Ensure the activity retention log is set to at least a year",
	Description: "The average time to detect a breach is up to 210 days, to ensure that all the information required for an effective investigation is available, the retention period should allow for delayed starts to investigating.",
	Impact:      "Short life activity logs can lead to missing records when investigating a breach",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/azure-monitor/essentials/platform-logs-overview", 
	},
}

