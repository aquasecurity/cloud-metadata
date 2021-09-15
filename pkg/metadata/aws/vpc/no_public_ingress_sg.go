package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngressSg = metadata.Metadata{
	ID:          "AVD-AWS-0106",
	Title:       "An inline ingress security group rule allows traffic from /0.",
	Description: "Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.",
	Impact:      "The port is exposed for ingress from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

