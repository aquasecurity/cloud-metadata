package database

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicFirewallAccess = metadata.Metadata{
	ID:          "AVD-AZU-0019",
	Title:       "Ensure database firewalls do not permit public access",
	Description: "Azure services can be allowed access through the firewall using a start and end IP address of 0.0.0.0. No other end ip address should be combined with a start of 0.0.0.0",
	Impact:      "Publicly accessible databases could lead to compromised data",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/rest/api/sql/2021-02-01-preview/firewall-rules/create-or-update", 
	},
}

