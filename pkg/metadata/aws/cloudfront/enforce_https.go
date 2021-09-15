package cloudfront

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforceHttps = metadata.Metadata{
	ID:          "AVD-AWS-0012",
	Title:       "CloudFront distribution allows unencrypted (HTTP) communications.",
	Description: "Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection, they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.",
	Impact:      "CloudFront is available through an unencrypted connection",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-cloudfront-to-s3-origin.html", 
	},
}

