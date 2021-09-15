package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var BlockPublicAcls = metadata.Metadata{
	ID:          "AVD-AWS-0086",
	Title:       "S3 Access block should block public ACL",
	Description: "S3 buckets should block public ACLs on buckets and any objects they contain. By blocking, PUTs with fail if the object has any public ACL a.",
	Impact:      "PUT calls with public ACLs specified can make objects public",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html", 
	},
}

