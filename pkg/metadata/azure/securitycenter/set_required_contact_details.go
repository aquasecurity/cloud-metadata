package securitycenter

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SetRequiredContactDetails = metadata.Metadata{
	ID:          "AVD-AZU-0046",
	Title:       "The required contact details should be set for security center",
	Description: "It is recommended that at least one valid contact is configured for the security center. 
Microsoft will notify the security contact directly in the event of a security incident and will look to use a telephone number in cases where a prompt response is required.",
	Impact:      "Without a telephone number set, Azure support can't contact",
	Severity:    "LOW",
	Links:       []string {
		"https://azure.microsoft.com/en-us/services/security-center/", 
	},
}

