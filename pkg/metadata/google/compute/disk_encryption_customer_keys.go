package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DiskEncryptionCustomerKeys = metadata.Metadata{
	ID:          "AVD-GCP-0003",
	Title:       "Encrypted compute disk with unmanaged keys.",
	Description: "By default, Compute Engine encrypts all data at rest. Compute Engine handles and manages this encryption for you without any additional actions on your part.

If the disk_encryption_key block is included in the resource declaration then it *must* include a raw_key or kms_key_self_link.",
	Impact:      "Encryption of disk using unmanaged keys.",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/compute/docs/disks/customer-supplied-encryption", 
	},
}

