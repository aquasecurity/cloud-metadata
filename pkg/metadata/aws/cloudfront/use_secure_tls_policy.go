package cloudfront

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-AWS-0013",
	Title:       "CloudFront distribution uses outdated SSL/TLS protocols.",
	Description: "You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.",
	Impact:      "Outdated SSL policies increase exposure to known vulnerabilities",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/secure-connections-supported-viewer-protocols-ciphers.html", 
	},
}

