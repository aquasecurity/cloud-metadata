package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EncryptInTransitData = metadata.Metadata{
	ID:          "AVD-GCP-0053",
	Title:       "SSL connections to a SQL database instance should be enforced.",
	Description: "In-transit data should be encrypted so that if traffic is intercepted data will not be exposed in plaintext to attackers.",
	Impact:      "Intercepted data can be read in transit",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/sql/docs/mysql/configure-ssl-instance", 
	},
}

