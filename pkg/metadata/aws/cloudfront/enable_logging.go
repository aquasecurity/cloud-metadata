package cloudfront

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableLogging = metadata.Metadata{
	ID:          "AVD-AWS-0010",
	Title:       "Cloudfront distribution should have Access Logging configured",
	Description: "You should configure CloudFront Access Logging to create log files that contain detailed information about every user request that CloudFront receives",
	Impact:      "Logging provides vital information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html", 
	},
}

