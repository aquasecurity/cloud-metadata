package keyvault

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnsureKeyExpiry = metadata.Metadata{
	ID:          "AVD-AZU-0029",
	Title:       "Ensure that the expiration date is set on all keys",
	Description: "Expiration Date is an optional Key Vault Key behavior and is not set by default.

Set when the resource will be become inactive.",
	Impact:      "Long life keys increase the attack surface when compromised",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.microsoft.com/en-us/powershell/module/az.keyvault/update-azkeyvaultkey?view=azps-5.8.0#example-1--modify-a-key-to-enable-it--and-set-the-expiration-date-and-tags", 
	},
}

