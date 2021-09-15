package authorization

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var LimitRoleActions = metadata.Metadata{
	ID:          "AVD-AZU-0007",
	Title:       "Roles limited to the required actions",
	Description: "The permissions granted to a role should be kept to the minimum required to be able to do the task. Wildcard permissions must not be used.",
	Impact:      "Open permissions for subscriptions could result in an easily compromisable account",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

