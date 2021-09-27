package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DiskEncryptionRequired = metadata.Metadata{
	ID:          "AVD-GCP-0004",
	Title:       "The encryption key used to encrypt a compute disk has been specified in plaintext.",
	Description: "Sensitive values such as raw encryption keys should not be included in your Terraform code, and should be stored securely by a secrets manager.",
	Impact:      "The encryption key should be considered compromised as it is not stored securely.",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://cloud.google.com/compute/docs/disks/customer-supplied-encryption", 
	},
}

