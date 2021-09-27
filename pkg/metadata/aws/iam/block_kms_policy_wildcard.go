package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var BlockKmsPolicyWildcard = metadata.Metadata{
	ID:          "AVD-AWS-0055",
	Title:       "IAM customer managed policies should not allow decryption actions on all KMS keys",
	Description: "IAM policies define which actions an identity (user, group, or role) can perform on which resources. Following security best practices, AWS recommends that you allow least privilege. In other words, you should grant to identities only the kms:Decrypt or kms:ReEncryptFrom permissions and only for the keys that are required to perform a task.",
	Impact:      "Identities may be able to decrypt data which they should not have access to",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html#fsbp-kms-1", 
	},
}

