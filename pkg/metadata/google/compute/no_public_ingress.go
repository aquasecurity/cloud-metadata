package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngress = metadata.Metadata{
	ID:          "AVD-GCP-0014",
	Title:       "An inbound firewall rule allows traffic from /0.",
	Description: "Network security rules should not use very broad subnets.

Where possible, segments should be broken into smaller subnets and avoid using the /0 subnet.",
	Impact:      "The port is exposed for ingress from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://cloud.google.com/vpc/docs/using-firewalls", 
	},
}

