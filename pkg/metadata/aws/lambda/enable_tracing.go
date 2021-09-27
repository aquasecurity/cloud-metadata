package lambda

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableTracing = metadata.Metadata{
	ID:          "AVD-AWS-0066",
	Title:       "Lambda functions should have X-Ray tracing enabled",
	Description: "X-Ray tracing enables end-to-end debugging and analysis of all function activity. This will allow for identifying bottlenecks, slow downs and timeouts.",
	Impact:      "WIthout full tracing enabled it is difficult to trace the flow of logs",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

