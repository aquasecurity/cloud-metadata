package keyvault

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SpecifyNetworkAcl = metadata.Metadata{
	ID:          "AVD-AZU-0032",
	Title:       "Key vault should have the network acl block specified",
	Description: "Network ACLs allow you to reduce your exposure to risk by limiting what can access your key vault. 

The default action of the Network ACL should be set to deny for when IPs are not matched. Azure services can be allowed to bypass.",
	Impact:      "Without a network ACL the key vault is freely accessible",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/key-vault/general/network-security", 
	},
}

