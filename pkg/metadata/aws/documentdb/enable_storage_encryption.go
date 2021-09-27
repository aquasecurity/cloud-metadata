package documentdb

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableStorageEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0021",
	Title:       "DocumentDB storage must be encrypted",
	Description: "Encryption of the underlying storage used by DocumentDB ensures that if their is compromise of the disks, the data is still protected.",
	Impact:      "Unencrypted sensitive data is vulnerable to compromise.",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

