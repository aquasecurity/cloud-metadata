package ssm

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SecretUseCustomerKey = metadata.Metadata{
	ID:          "AVD-AWS-0098",
	Title:       "Secrets Manager should use customer managed keys",
	Description: "Secrets Manager encrypts secrets by default using a default key created by AWS. To ensure control and granularity of secret encryption, CMK's should be used explicitly.",
	Impact:      "Using AWS managed keys reduces the flexibility and control over the encryption key",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.aws.amazon.com/kms/latest/developerguide/services-secrets-manager.html#asm-encrypt", 
	},
}

