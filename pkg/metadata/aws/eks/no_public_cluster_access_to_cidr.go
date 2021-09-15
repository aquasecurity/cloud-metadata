package eks

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicClusterAccessToCidr = metadata.Metadata{
	ID:          "AVD-AWS-0041",
	Title:       "EKS cluster should not have open CIDR range for public access",
	Description: "EKS Clusters have public access cidrs set to 0.0.0.0/0 by default which is wide open to the internet. This should be explicitly set to a more specific CIDR range",
	Impact:      "EKS can be access from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/eks/latest/userguide/create-public-private-vpc.html", 
	},
}

