package cloudwatch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var LogGroupCustomerKey = metadata.Metadata{
	ID:          "AVD-AWS-0017",
	Title:       "CloudWatch log groups should be encrypted using CMK",
	Description: "CloudWatch log groups are encrypted by default, however, to get the full benefit of controlling key rotation and other KMS aspects a KMS CMK should be used.",
	Impact:      "Log data may be leaked if the logs are compromised. No auditing of who have viewed the logs.",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/encrypt-log-data-kms.html", 
	},
}

