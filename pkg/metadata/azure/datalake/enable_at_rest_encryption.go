package datalake

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAtRestEncryption = metadata.Metadata{
	ID:          "AVD-AZU-0026",
	Title:       "Unencrypted data lake storage.",
	Description: "Datalake storage encryption defaults to Enabled, it shouldn't be overridden to Disabled.",
	Impact:      "Data could be read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/data-lake-store/data-lake-store-security-overview", 
	},
}

