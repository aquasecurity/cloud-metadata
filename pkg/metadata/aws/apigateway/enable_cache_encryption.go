package apigateway

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableCacheEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0002",
	Title:       "API Gateway must have cache enabled",
	Description: "Method cache encryption ensures that any sensitive data in the cache is not vulnerable to compromise in the event of interception",
	Impact:      "Data stored in the cache that is unencrypted may be vulnerable to compromise",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

