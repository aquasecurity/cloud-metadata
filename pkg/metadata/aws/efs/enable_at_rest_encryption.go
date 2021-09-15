package efs

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAtRestEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0037",
	Title:       "EFS Encryption has not been enabled",
	Description: "If your organization is subject to corporate or regulatory policies that require encryption of data and metadata at rest, we recommend creating a file system that is encrypted at rest, and mounting your file system using encryption of data in transit.",
	Impact:      "Data can be read from the EFS if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/efs/latest/ug/encryption.html", 
	},
}

