package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-AZU-0024",
	Title:       "Databases should have the minimum TLS set for connections",
	Description: "You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.",
	Impact:      "Outdated TLS policies increase exposure to known issues",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

