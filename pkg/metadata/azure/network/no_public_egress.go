package network

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicEgress = metadata.Metadata{
	ID:          "AVD-AZU-0040",
	Title:       "An outbound network security rule allows traffic to /0.",
	Description: "Network security rules should not use very broad subnets.

Where possible, segments should be broken into smaller subnets.",
	Impact:      "The port is exposed for egress to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/security/fundamentals/network-best-practices", 
	},
}

