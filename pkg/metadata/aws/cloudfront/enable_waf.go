package cloudfront

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableWaf = metadata.Metadata{
	ID:          "AVD-AWS-0011",
	Title:       "CloudFront distribution does not have a WAF in front.",
	Description: "You should configure a Web Application Firewall in front of your CloudFront distribution. This will mitigate many types of attacks on your web application.",
	Impact:      "Complex web application attacks can more easily be performed without a WAF",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/waf/latest/developerguide/cloudfront-features.html", 
	},
}

