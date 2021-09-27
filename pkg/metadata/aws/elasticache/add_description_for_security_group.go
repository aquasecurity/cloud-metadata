package elasticache

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AddDescriptionForSecurityGroup = metadata.Metadata{
	ID:          "AVD-AWS-0049",
	Title:       "Missing description for security group/security group rule.",
	Description: "Security groups and security group rules should include a description for auditing purposes.

Simplifies auditing, debugging, and managing security groups.",
	Impact:      "Descriptions provide context for the firewall rule reasons",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

