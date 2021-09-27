package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIp = metadata.Metadata{
	ID:          "AVD-OCI-0001",
	Title:       "Compute instance requests an IP reservation from a public pool",
	Description: "Compute instance requests an IP reservation from a public pool

The compute instance has the ability to be reached from outside, you might want to sonder the use of a non public IP.",
	Impact:      "The compute instance has the ability to be reached from outside",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

