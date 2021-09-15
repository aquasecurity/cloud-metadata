package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableVpcFlowLogs = metadata.Metadata{
	ID:          "AVD-GCP-0006",
	Title:       "VPC flow logs should be enabled for all subnets",
	Description: "VPC flow logs record information about all traffic, which is a vital tool in reviewing anomalous traffic.",
	Impact:      "Limited auditing capability and awareness",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

