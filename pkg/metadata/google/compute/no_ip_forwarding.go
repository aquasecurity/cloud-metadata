package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoIpForwarding = metadata.Metadata{
	ID:          "AVD-GCP-0008",
	Title:       "Instances should not have IP forwarding enabled",
	Description: "Disabling IP forwarding ensuresthe instance can only receive packets addressed to the instance and can only send packets with a source address of the instance.",
	Impact:      "Instance can send/receive packets without the explicit instance address",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

