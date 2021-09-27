package ecs

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableInTransitEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0035",
	Title:       "ECS Task Definitions with EFS volumes should use in-transit encryption",
	Description: "ECS task definitions that have volumes using EFS configuration should explicitly enable in transit encryption to prevent the risk of data loss due to interception.",
	Impact:      "Intercepted traffic to and from EFS may lead to data loss",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonECS/latest/userguide/efs-volumes.html", "https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html", 
	},
}

