package securitycenter

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableStandardSubscription = metadata.Metadata{
	ID:          "AVD-AZU-0045",
	Title:       "Enable the standard security center subscription tier",
	Description: "To benefit from Azure Defender you should use the Standard subscription tier.
			
			Enabling Azure Defender extends the capabilities of the free mode to workloads running in private and other public clouds, providing unified security management and threat protection across your hybrid cloud workloads.",
	Impact:      "Using free subscription does not enable Azure Defender for the resource type",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/security-center/security-center-pricing", 
	},
}

