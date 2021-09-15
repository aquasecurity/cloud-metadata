package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoDefaultServiceAccount = metadata.Metadata{
	ID:          "AVD-GCP-0007",
	Title:       "Instances should not use the default service account",
	Description: "The default service account has full project access. Instances should instead be assigned the minimal access they need.",
	Impact:      "Instance has full access to the project",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

