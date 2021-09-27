package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AddDescriptionToSecurityGroup = metadata.Metadata{
	ID:          "AVD-AWS-0099",
	Title:       "Missing description for security group/security group rule.",
	Description: "Security groups and security group rules should include a description for auditing purposes.

Simplifies auditing, debugging, and managing security groups.",
	Impact:      "Descriptions provide context for the firewall rule reasons",
	Severity:    "LOW",
	Links:       []string {
		"https://www.cloudconformity.com/knowledge-base/aws/EC2/security-group-rules-description.html", 
	},
}

