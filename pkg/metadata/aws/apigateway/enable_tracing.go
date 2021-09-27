package apigateway

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableTracing = metadata.Metadata{
	ID:          "AVD-AWS-0003",
	Title:       "API Gateway must have X-Ray tracing enabled",
	Description: "X-Ray tracing enables end-to-end debugging and analysis of all API Gateway HTTP requests.",
	Impact:      "WIthout full tracing enabled it is difficult to trace the flow of logs",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.aws.amazon.com/xray/latest/devguide/xray-services-apigateway.html", 
	},
}

