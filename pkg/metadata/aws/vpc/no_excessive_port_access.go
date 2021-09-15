package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoExcessivePortAccess = metadata.Metadata{
	ID:          "AVD-AWS-0102",
	Title:       "An ingress Network ACL rule allows ALL ports.",
	Description: "Ensure access to specific required ports is allowed, and nothing else.",
	Impact:      "All ports exposed for egressing data",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html", 
	},
}

