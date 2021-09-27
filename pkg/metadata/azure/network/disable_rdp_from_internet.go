package network

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DisableRdpFromInternet = metadata.Metadata{
	ID:          "AVD-AZU-0039",
	Title:       "RDP access should not be accessible from the Internet, should be blocked on port 3389",
	Description: "RDP access can be configured on either the network security group or in the network security group rule.

RDP access should not be permitted from the internet (*, 0.0.0.0, /0, internet, any). Consider using the Azure Bastion Service.",
	Impact:      "Anyone from the internet can potentially RDP onto an instance",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/bastion/tutorial-create-host-portal", 
	},
}

