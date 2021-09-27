package apigateway

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAccessLogging = metadata.Metadata{
	ID:          "AVD-AWS-0001",
	Title:       "API Gateway stages for V1 and V2 should have access logging enabled",
	Description: "API Gateway stages should have access log settings block configured to track all access to a particular stage. This should be applied to both v1 and v2 gateway stages.",
	Impact:      "Logging provides vital information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html", 
	},
}

