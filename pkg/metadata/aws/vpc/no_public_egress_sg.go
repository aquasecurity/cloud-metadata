package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicEgressSg = metadata.Metadata{
	ID:          "AVD-AWS-0103",
	Title:       "An inline egress security group rule allows traffic to /0.",
	Description: "Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.",
	Impact:      "The port is exposed for egressing data to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

