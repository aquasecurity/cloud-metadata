package ecr

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-AWS-0032",
	Title:       "ECR repository policy must block public access",
	Description: "Allowing public access to the ECR repository risks leaking sensitive of abusable information",
	Impact:      "Risk of potential data leakage of sensitive artifacts",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

