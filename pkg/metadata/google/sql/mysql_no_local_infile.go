package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var MysqlNoLocalInfile = metadata.Metadata{
	ID:          "AVD-GCP-0054",
	Title:       "Disable local_infile setting in MySQL",
	Description: "Arbitrary files can be read from the system using LOAD_DATA unless this setting is disabled.",
	Impact:      "Arbitrary files read by attackers when combined with a SQL injection vulnerability.",
	Severity:    "HIGH",
	Links:       []string {
		"https://dev.mysql.com/doc/refman/8.0/en/load-data-local-security.html", 
	},
}

