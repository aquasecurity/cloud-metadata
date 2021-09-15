package cloudtrail

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableLogValidation = metadata.Metadata{
	ID:          "AVD-AWS-0016",
	Title:       "Cloudtrail log validation should be enabled to prevent tampering of log data",
	Description: "Log validation should be activated on Cloudtrail logs to prevent the tampering of the underlying data in the S3 bucket. It is feasible that a rogue actor compromising an AWS account might want to modify the log data to remove trace of their actions.",
	Impact:      "Illicit activity could be removed from the logs",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-log-file-validation-intro.html", 
	},
}

