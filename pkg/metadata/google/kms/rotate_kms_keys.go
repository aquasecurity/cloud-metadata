package kms

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RotateKmsKeys = metadata.Metadata{
	ID:          "AVD-GCP-0049",
	Title:       "KMS keys should be rotated at least every 90 days",
	Description: "Keys should be rotated on a regular basis to limit exposure if a given key should become compromised.",
	Impact:      "Exposure is greater if the same keys are used over a long period",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

