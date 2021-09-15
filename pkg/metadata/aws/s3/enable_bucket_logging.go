package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableBucketLogging = metadata.Metadata{
	ID:          "AVD-AWS-0089",
	Title:       "S3 Bucket does not have logging enabled.",
	Description: "Buckets should have logging enabled so that access can be audited.",
	Impact:      "There is no way to determine the access to this bucket",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html", 
	},
}

