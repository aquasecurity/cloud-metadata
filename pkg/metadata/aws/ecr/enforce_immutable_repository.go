package ecr

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforceImmutableRepository = metadata.Metadata{
	ID:          "AVD-AWS-0031",
	Title:       "ECR images tags shouldn't be mutable.",
	Description: "ECR images should be set to IMMUTABLE to prevent code injection through image mutation.

This can be done by setting image_tab_mutability to IMMUTABLE",
	Impact:      "Image tags could be overwritten with compromised images",
	Severity:    "HIGH",
	Links:       []string {
		"https://sysdig.com/blog/toctou-tag-mutability/", 
	},
}

