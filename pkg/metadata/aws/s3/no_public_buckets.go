package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicBuckets = metadata.Metadata{
	ID:          "AVD-AWS-0093",
	Title:       "S3 Access block should restrict public bucket to limit access",
	Description: "S3 buckets should restrict public policies for the bucket. By enabling, the restrict_public_buckets, only the bucket owner and AWS Services can access if it has a public policy.",
	Impact:      "Public buckets can be accessed by anyone",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/dev-retired/access-control-block-public-access.html", 
	},
}

