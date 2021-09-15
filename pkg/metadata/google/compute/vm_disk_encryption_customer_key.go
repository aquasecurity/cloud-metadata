package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var VmDiskEncryptionCustomerKey = metadata.Metadata{
	ID:          "AVD-GCP-0019",
	Title:       "VM disks should be encrypted with Customer Supplied Encryption Keys",
	Description: "Using unmanaged keys makes rotation and general management difficult.",
	Impact:      "Using unmanaged keys does not allow for proper management",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

