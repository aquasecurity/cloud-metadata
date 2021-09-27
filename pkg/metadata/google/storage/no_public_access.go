package storage

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-GCP-0065",
	Title:       "Ensure that Cloud Storage bucket is not anonymously or publicly accessible.",
	Description: "Using 'allUsers' or 'allAuthenticatedUsers' as members in an IAM member/binding causes data to be exposed outside of the organisation.",
	Impact:      "Public exposure of sensitive data.",
	Severity:    "HIGH",
	Links:       []string {
		"https://jbrojbrojbro.medium.com/you-make-the-rules-with-authentication-controls-for-cloud-storage-53c32543747b", 
	},
}

