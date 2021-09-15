package eks

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicClusterAccess = metadata.Metadata{
	ID:          "AVD-AWS-0040",
	Title:       "EKS Clusters should have the public access disabled",
	Description: "EKS clusters are available publicly by default, this should be explicitly disabled in the vpc_config of the EKS cluster resource.",
	Impact:      "EKS can be access from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html", 
	},
}

