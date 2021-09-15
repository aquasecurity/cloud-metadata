package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableSslEnforcement = metadata.Metadata{
	ID:          "AVD-AZU-0017",
	Title:       "SSL should be enforced on database connections where applicable",
	Description: "SSL connections should be enforced were available to ensure secure transfer and reduce the risk of compromising data in flight.",
	Impact:      "Insecure connections could lead to data loss and other vulnerabilities",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

