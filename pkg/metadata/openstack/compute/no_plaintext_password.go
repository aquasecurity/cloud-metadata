package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPlaintextPassword = metadata.Metadata{
	ID:          "AVD-OSK-0001",
	Title:       "No plaintext password for compute instance",
	Description: "Assigning a password to the compute instance using plaintext could lead to compromise; it would be preferable to use key-pairs as a login mechanism",
	Impact:      "Including a plaintext password could lead to compromised instance",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

