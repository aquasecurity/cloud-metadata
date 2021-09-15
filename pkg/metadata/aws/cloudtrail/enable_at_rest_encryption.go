package cloudtrail

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAtRestEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0015",
	Title:       "Cloudtrail should be encrypted at rest to secure access to sensitive trail data",
	Description: "Cloudtrail logs should be encrypted at rest to secure the sensitive data. Cloudtrail logs record all activity that occurs in the the account through API calls and would be one of the first places to look when reacting to a breach.",
	Impact:      "Data can be freely read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/awscloudtrail/latest/userguide/encrypting-cloudtrail-log-files-with-aws-kms.html", 
	},
}

