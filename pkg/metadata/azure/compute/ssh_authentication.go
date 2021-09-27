package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SshAuthentication = metadata.Metadata{
	ID:          "AVD-AZU-0011",
	Title:       "Password authentication in use instead of SSH keys.",
	Description: "Access to instances should be authenticated using SSH keys. Removing the option of password authentication enforces more secure methods while removing the risks inherent with passwords.",
	Impact:      "Passwords are potentially easier to compromise than SSH Keys",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/virtual-machines/linux/create-ssh-keys-detailed", 
	},
}

