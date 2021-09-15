package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPlaintextVmDiskKeys = metadata.Metadata{
	ID:          "AVD-GCP-0011",
	Title:       "VM disk encryption keys should not be provided in plaintext",
	Description: "Providing your encryption key in plaintext format means anyone with access to the source code also has access to the key.",
	Impact:      "Compromise of encryption keys",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

