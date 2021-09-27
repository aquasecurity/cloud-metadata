package mq

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableGeneralLogging = metadata.Metadata{
	ID:          "AVD-AWS-0071",
	Title:       "MQ Broker should have general logging enabled",
	Description: "Logging should be enabled to allow tracing of issues and activity to be investigated more fully. Logs provide additional information and context which is often invalauble during investigation",
	Impact:      "Without logging it is difficult to trace issues",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

