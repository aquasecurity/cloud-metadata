package network

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngress = metadata.Metadata{
	ID:          "AVD-AZU-0041",
	Title:       "An inbound network security rule allows traffic from /0.",
	Description: "Network security rules should not use very broad subnets.

Where possible, segments should be broken into smaller subnets.",
	Impact:      "The port is exposed for ingress from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/security/fundamentals/network-best-practices", 
	},
}

