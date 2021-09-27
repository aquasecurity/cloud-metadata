package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableStackdriverMonitoring = metadata.Metadata{
	ID:          "AVD-GCP-0029",
	Title:       "Stackdriver Monitoring should be enabled",
	Description: "StackDriver monitoring aggregates logs, events, and metrics from your Kubernetes environment on GKE to help you understand your application's behavior in production.",
	Impact:      "Visibility will be reduced",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

