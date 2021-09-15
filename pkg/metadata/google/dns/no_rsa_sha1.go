package dns

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoRsaSha1 = metadata.Metadata{
	ID:          "AVD-GCP-0021",
	Title:       "Zone signing should not use RSA SHA1",
	Description: "RSA SHA1 is a weaker algorithm than SHA2-based algorithms such as RSA SHA256/512",
	Impact:      "Less secure encryption algorithm than others available",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

