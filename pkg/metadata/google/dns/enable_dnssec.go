package dns

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableDnssec = metadata.Metadata{
	ID:          "AVD-GCP-0020",
	Title:       "Cloud DNS should use DNSSEC",
	Description: "DNSSEC authenticates DNS responses, preventing MITM attacks and impersonation.",
	Impact:      "Unverified DNS responses could lead to man-in-the-middle attacks",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

