package keyvault

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPurge = metadata.Metadata{
	ID:          "AVD-AZU-0031",
	Title:       "Key vault should have purge protection enabled",
	Description: "Purge protection is an optional Key Vault behavior and is not enabled by default.

Purge protection can only be enabled once soft-delete is enabled. It can be turned on via CLI or PowerShell.",
	Impact:      "Keys could be purged from the vault without protection",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview#purge-protection", 
	},
}

