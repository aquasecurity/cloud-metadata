package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SpecifyPublicAccessBlock = metadata.Metadata{
	ID:          "AVD-AWS-0094",
	Title:       "S3 buckets should each define an aws_s3_bucket_public_access_block",
	Description: "The "block public access" settings in S3 override individual policies that apply to a given bucket, meaning that all public access can be controlled in one central definition for that bucket. It is therefore good practice to define these settings for each bucket in order to clearly define the public access that can be allowed for it.",
	Impact:      "Public access policies may be applied to sensitive data buckets",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html", 
	},
}

