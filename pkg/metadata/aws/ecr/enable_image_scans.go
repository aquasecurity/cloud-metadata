package ecr

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableImageScans = metadata.Metadata{
	ID:          "AVD-AWS-0030",
	Title:       "ECR repository has image scans disabled.",
	Description: "Repository image scans should be enabled to ensure vulnerable software can be discovered and remediated as soon as possible.",
	Impact:      "The ability to scan images is not being used and vulnerabilities will not be highlighted",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html", 
	},
}

