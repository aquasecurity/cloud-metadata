package mq

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-AWS-0072",
	Title:       "Ensure MQ Broker is not publicly exposed",
	Description: "Public access of the MQ broker should be disabled and only allow routes to applications that require access.",
	Impact:      "Publicly accessible MQ Broker may be vulnerable to compromise",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

