package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DisallowMixedSgr = metadata.Metadata{
	ID:          "AVD-AWS-0100",
	Title:       "Ensures that usage of security groups with inline rules and security group rule resources are not mixed.",
	Description: "Mixing Terraform standalone security_group_rule resource and security_group resource with inline ingress/egress rules results in rules being overwritten during Terraform apply.",
	Impact:      "Security group rules will be overwritten and will result in unintended blocking of network traffic",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

