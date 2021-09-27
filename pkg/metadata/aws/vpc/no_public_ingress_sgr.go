package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngressSgr = metadata.Metadata{
	ID:          "AVD-AWS-0107",
	Title:       "An ingress security group rule allows traffic from /0.",
	Description: "Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.",
	Impact:      "Your port exposed to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html", 
	},
}

