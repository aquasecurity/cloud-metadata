package datafactory

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-AZU-0025",
	Title:       "Data Factory should have public access disabled, the default is enabled.",
	Description: "Data Factory has public access set to true by default.

Disabling public network access is applicable only to the self-hosted integration runtime, not to Azure Integration Runtime and SQL Server Integration Services (SSIS) Integration Runtime.",
	Impact:      "Data factory is publicly accessible",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/data-factory/data-movement-security-considerations#hybrid-scenarios", 
	},
}

