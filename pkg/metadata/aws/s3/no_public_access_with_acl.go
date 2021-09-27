package s3

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccessWithAcl = metadata.Metadata{
	ID:          "AVD-AWS-0092",
	Title:       "S3 Bucket has an ACL defined which allows public access.",
	Description: "S3 bucket permissions should be set to deny public access unless explicitly required.

Granting write access publicly with public-read-write is especially dangerous as you will be billed for any uploaded files.

Additionally, you should not use the authenticated-read canned ACL, as this provides read access to any authenticated AWS user, not just AWS users within your organisation.",
	Impact:      "The contents of the bucket can be accessed publicly",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://aws.amazon.com/premiumsupport/knowledge-center/secure-s3-resources/", 
	},
}

