package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableStackdriverLogging = metadata.Metadata{
	ID:          "AVD-GCP-0028",
	Title:       "Stackdriver Logging should be enabled",
	Description: "StackDriver logging provides a useful interface to all of stdout/stderr for each container and should be enabled for moitoring, debugging, etc.",
	Impact:      "Visibility will be reduced",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

