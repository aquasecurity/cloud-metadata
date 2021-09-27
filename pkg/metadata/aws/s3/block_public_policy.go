package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var BlockPublicPolicy = metadata.Metadata{
	ID:          "AVD-AWS-0087",
	Title:       "S3 Access block should block public policy",
	Description: "S3 bucket policy should have block public policy to prevent users from putting a policy that enable public access.",
	Impact:      "Users could put a policy that allows public access",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/dev-retired/access-control-block-public-access.html", 
	},
}

