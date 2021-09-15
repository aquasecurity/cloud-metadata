package appservice

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableHttp2 = metadata.Metadata{
	ID:          "AVD-AZU-0003",
	Title:       "Web App uses the latest HTTP version",
	Description: "Use the latest version of HTTP to ensure you are benefiting from security fixes",
	Impact:      "Outdated versions of HTTP has security vulnerabilities",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

