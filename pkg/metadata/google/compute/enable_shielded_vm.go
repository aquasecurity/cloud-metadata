package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableShieldedVm = metadata.Metadata{
	ID:          "AVD-GCP-0005",
	Title:       "Instances should have Shielded VM enabled",
	Description: "A Shielded VM is a VM with enhanced defences/detection for rootkits/bootkits.",
	Impact:      "Unable to detect rootkits",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

