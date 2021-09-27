package rds

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoClassicResources = metadata.Metadata{
	ID:          "AVD-AWS-0081",
	Title:       "AWS Classic resource usage.",
	Description: "AWS Classic resources run in a shared environment with infrastructure owned by other AWS customers. You should run
resources in a VPC instead.",
	Impact:      "Classic resources are running in a shared environment with other customers",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-classic-platform.html", 
	},
}

