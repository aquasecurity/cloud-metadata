package secrets

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SensitiveInAttributeValue = metadata.Metadata{
	ID:          "AVD-GEN-0002",
	Title:       "The attribute has potentially sensitive data, passwords, tokens or keys in it",
	Description: "Sensitive data stored in attributes can result in compromised data. Sensitive data should be passed in through secret variables",
	Impact:      "Sensitive credentials may be compromised",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

