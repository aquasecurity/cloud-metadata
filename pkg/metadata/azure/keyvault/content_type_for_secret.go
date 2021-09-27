package keyvault

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var ContentTypeForSecret = metadata.Metadata{
	ID:          "AVD-AZU-0028",
	Title:       "Key vault Secret should have a content type set",
	Description: "Content Type is an optional Key Vault Secret behavior and is not enabled by default.

Clients may specify the content type of a secret to assist in interpreting the secret data when it's retrieved. The maximum length of this field is 255 characters. There are no pre-defined values. The suggested usage is as a hint for interpreting the secret data.",
	Impact:      "The secret's type is unclear without a content type",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/key-vault/secrets/about-secrets", 
	},
}

