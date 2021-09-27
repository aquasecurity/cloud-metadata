package neptune

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableStorageEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0076",
	Title:       "Neptune storage must be encrypted at rest",
	Description: "Encryption of Neptune storage ensures that if their is compromise of the disks, the data is still protected.",
	Impact:      "Unencrypted sensitive data is vulnerable to compromise.",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

