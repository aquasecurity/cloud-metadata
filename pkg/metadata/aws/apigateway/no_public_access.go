package apigateway

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-AWS-0004",
	Title:       "No public access to API Gateway methods",
	Description: "API Gateway methods should be protected by authorization or api key. OPTION verb calls can be used without authorization",
	Impact:      "API gateway methods can be unauthorized accessed",
	Severity:    "LOW",
	Links:       []string {
		"https://aws.amazon.com/blogs/compute/introducing-amazon-api-gateway-private-endpoints", 
	},
}

