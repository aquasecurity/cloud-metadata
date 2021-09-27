package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-GCP-0057",
	Title:       "Ensure that Cloud SQL Database Instances are not publicly exposed",
	Description: "Database instances should be configured so that they are not available over the public internet, but to internal compute resources which access them.",
	Impact:      "Public exposure of sensitive data",
	Severity:    "HIGH",
	Links:       []string {
		"https://www.cloudconformity.com/knowledge-base/gcp/CloudSQL/publicly-accessible-cloud-sql-instances.html", 
	},
}

