package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngress = metadata.Metadata{
	ID:          "AVD-AWS-0105",
	Title:       "An ingress Network ACL rule allows specific ports from /0.",
	Description: "Opening up ACLs to the public internet is potentially dangerous. You should restrict access to IP addresses or ranges that explicitly require it where possible.",
	Impact:      "The ports are exposed for ingressing data to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html", 
	},
}

