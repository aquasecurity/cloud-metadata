package msk

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableLogging = metadata.Metadata{
	ID:          "AVD-AWS-0074",
	Title:       "Ensure MSK Cluster logging is enabled",
	Description: "Managed streaming for Kafka can log to Cloud Watch, Kinesis Firehose and S3, at least one of these locations should be logged to",
	Impact:      "Without logging it is difficult to trace issues",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

