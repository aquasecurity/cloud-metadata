package sns

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableTopicEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0095",
	Title:       "Unencrypted SNS topic.",
	Description: "Queues should be encrypted with customer managed KMS keys and not default AWS managed keys, in order to allow granular control over access to specific queues.",
	Impact:      "The SNS topic messages could be read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html", 
	},
}

