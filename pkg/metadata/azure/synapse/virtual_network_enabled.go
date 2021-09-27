package synapse

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var VirtualNetworkEnabled = metadata.Metadata{
	ID:          "AVD-AZU-0053",
	Title:       "Synapse Workspace should have managed virtual network enabled, the default is disabled.",
	Description: "Synapse Workspace does not have managed virtual network enabled by default.

When you create your Azure Synapse workspace, you can choose to associate it to a Microsoft Azure Virtual Network. The Virtual Network associated with your workspace is managed by Azure Synapse. This Virtual Network is called a Managed workspace Virtual Network.
Managed private endpoints are private endpoints created in a Managed Virtual Network associated with your Azure Synapse workspace. Managed private endpoints establish a private link to Azure resources. You can only use private links in a workspace that has a Managed workspace Virtual Network.",
	Impact:      "Your Synapse workspace is not using the private endpoints",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/synapse-analytics/security/synapse-workspace-managed-private-endpoints", "https://docs.microsoft.com/en-us/azure/synapse-analytics/security/synapse-workspace-managed-vnet", 
	},
}

