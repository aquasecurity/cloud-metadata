package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicEgress = metadata.Metadata{
	ID:          "AVD-GCP-0013",
	Title:       "An outbound firewall rule allows traffic to /0.",
	Description: "Network security rules should not use very broad subnets.

Where possible, segments should be broken into smaller subnets and avoid using the /0 subnet.",
	Impact:      "The port is exposed for egress to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://cloud.google.com/vpc/docs/using-firewalls", 
	},
}

