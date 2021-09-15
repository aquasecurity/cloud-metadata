package eks

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptSecrets = metadata.Metadata{
	ID:          "AVD-AWS-0039",
	Title:       "EKS should have the encryption of secrets enabled",
	Description: "EKS cluster resources should have the encryption_config block set with protection of the secrets resource.",
	Impact:      "EKS secrets could be read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-eks-adds-envelope-encryption-for-secrets-with-aws-kms/", 
	},
}

