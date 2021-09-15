package monitor

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var CaptureAllRegions = metadata.Metadata{
	ID:          "AVD-AZU-0035",
	Title:       "Ensure activitys are captured for all locations",
	Description: "Log profiles should capture all regions to ensure that all events are logged",
	Impact:      "Activity may be occurring in locations that aren't being monitored",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/cli/azure/monitor/log-profiles?view=azure-cli-latest#az_monitor_log_profiles_create-required-parameters", 
	},
}

