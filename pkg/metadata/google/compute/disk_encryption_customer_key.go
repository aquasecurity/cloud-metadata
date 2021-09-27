package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DiskEncryptionCustomerKey = metadata.Metadata{
	ID:          "AVD-GCP-0002",
	Title:       "Disks should be encrypted with Customer Supplied Encryption Keys",
	Description: "Using unmanaged keys makes rotation and general management difficult.",
	Impact:      "Using unmanaged keys does not allow for proper management",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

