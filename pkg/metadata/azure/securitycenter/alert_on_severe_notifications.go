package securitycenter

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AlertOnSevereNotifications = metadata.Metadata{
	ID:          "AVD-AZU-0044",
	Title:       "Send notification emails for high severity alerts",
	Description: "It is recommended that at least one valid contact is configured for the security center. 
Microsoft will notify the security contact directly in the event of a security incident using email and require alerting to be turned on.",
	Impact:      "The ability to react to high severity notifications could be delayed",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://azure.microsoft.com/en-us/services/security-center/", 
	},
}

