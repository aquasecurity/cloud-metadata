package monitor

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var CaptureAllActivities = metadata.Metadata{
	ID:          "AVD-AZU-0034",
	Title:       "Ensure log profile captures all activities",
	Description: "Log profiles should capture all categories to ensure that all events are logged",
	Impact:      "Log profile must capture all activity to be able to ensure that all relevant information possible is available for an investigation",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log", "https://docs.microsoft.com/en-us/cli/azure/monitor/log-profiles?view=azure-cli-latest#az_monitor_log_profiles_create-required-parameters", 
	},
}

