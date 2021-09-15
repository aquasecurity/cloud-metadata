package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableDiskEncryption = metadata.Metadata{
	ID:          "AVD-AZU-0009",
	Title:       "Enable disk encryption on managed disk",
	Description: "Manage disks should be encrypted at rest. When specifying the encryption_settings block, the enabled attribute should be set to true.",
	Impact:      "Data could be read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/virtual-machines/linux/disk-encryption", 
	},
}

