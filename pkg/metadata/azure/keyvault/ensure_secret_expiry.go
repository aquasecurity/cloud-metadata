package keyvault

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnsureSecretExpiry = metadata.Metadata{
	ID:          "AVD-AZU-0030",
	Title:       "Key Vault Secret should have an expiration date set",
	Description: "Expiration Date is an optional Key Vault Secret behavior and is not set by default.

Set when the resource will be become inactive.",
	Impact:      "Long life secrets increase the opportunity for compromise",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/key-vault/secrets/about-secrets", 
	},
}

