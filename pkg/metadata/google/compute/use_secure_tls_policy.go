package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-GCP-0018",
	Title:       "SSL policies should enforce secure versions of TLS",
	Description: "TLS versions prior to 1.2 are outdated and insecure. You should use 1.2 as aminimum version.",
	Impact:      "Data in transit is not sufficiently secured",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

