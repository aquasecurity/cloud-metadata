package storage

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DefaultActionDeny = metadata.Metadata{
	ID:          "AVD-AZU-0048",
	Title:       "The default action on Storage account network rules should be set to deny",
	Description: "The default_action for network rules should come into effect when no other rules are matched.

The default action should be set to Deny.",
	Impact:      "Network rules that allow could cause data to be exposed publicly",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/firewall/rule-processing", 
	},
}

