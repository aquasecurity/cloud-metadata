package mq

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAuditLogging = metadata.Metadata{
	ID:          "AVD-AWS-0070",
	Title:       "MQ Broker should have audit logging enabled",
	Description: "Logging should be enabled to allow tracing of issues and activity to be investigated more fully. Logs provide additional information and context which is often invalauble during investigation",
	Impact:      "Without audit logging it is difficult to trace activity in the MQ broker",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

