package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicEgressSgr = metadata.Metadata{
	ID:          "AVD-AWS-0104",
	Title:       "An egress security group rule allows traffic to /0.",
	Description: "Opening up ports to connect out to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that are explicitly required where possible.",
	Impact:      "Your port is egressing data to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

