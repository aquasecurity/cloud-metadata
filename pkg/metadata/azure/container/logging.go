package container

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var Logging = metadata.Metadata{
	ID:          "AVD-AZU-0014",
	Title:       "Ensure AKS logging to Azure Monitoring is Configured",
	Description: "Ensure AKS logging to Azure Monitoring is configured for containers to monitor the performance of workloads.",
	Impact:      "Logging provides valuable information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-onboard", 
	},
}

