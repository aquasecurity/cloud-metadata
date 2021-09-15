package storage

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var QueueServicesLoggingEnabled = metadata.Metadata{
	ID:          "AVD-AZU-0051",
	Title:       "When using Queue Services for a storage account, logging should be enabled.",
	Description: "Storage Analytics logs detailed information about successful and failed requests to a storage service. 

This information can be used to monitor individual requests and to diagnose issues with a storage service. 

Requests are logged on a best-effort basis.",
	Impact:      "Logging provides valuable information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/storage/common/storage-analytics-logging?tabs=dotnet", 
	},
}

