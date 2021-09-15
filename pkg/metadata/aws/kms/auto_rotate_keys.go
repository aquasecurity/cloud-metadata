package kms

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AutoRotateKeys = metadata.Metadata{
	ID:          "AVD-AWS-0065",
	Title:       "A KMS key is not configured to auto-rotate.",
	Description: "You should configure your KMS keys to auto rotate to maintain security and defend against compromise.",
	Impact:      "Long life KMS keys increase the attack surface when compromised",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html", 
	},
}

