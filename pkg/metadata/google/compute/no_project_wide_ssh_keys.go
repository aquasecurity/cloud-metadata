package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoProjectWideSshKeys = metadata.Metadata{
	ID:          "AVD-GCP-0012",
	Title:       "Disable project-wide SSH keys for all instances",
	Description: "Use of project-wide SSH keys means that a compromise of any one of these key pairs can result in all instances being compromised. It is recommended to use instance-level keys.",
	Impact:      "Compromise of a single key pair compromises all instances",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

