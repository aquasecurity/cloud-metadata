package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPlaintextDiskKeys = metadata.Metadata{
	ID:          "AVD-GCP-0010",
	Title:       "Disk encryption keys should not be provided in plaintext",
	Description: "Providing your encryption key in plaintext format means anyone with access to the source code also has access to the key.",
	Impact:      "Compromise of encryption keys",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

