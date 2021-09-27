package bigquery

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-GCP-0001",
	Title:       "BigQuery datasets should only be accessible within the organisation",
	Description: "Using 'allAuthenticatedUsers' provides any GCP user - even those outside of your organisation - access to your BigQuery dataset.",
	Impact:      "Exposure of sensitive data to the public iniernet",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

