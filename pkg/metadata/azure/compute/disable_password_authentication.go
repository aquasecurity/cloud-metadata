package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DisablePasswordAuthentication = metadata.Metadata{
	ID:          "AVD-AZU-0008",
	Title:       "Password authentication should be disabled on Azure virtual machines",
	Description: "Access to virtual machines should be authenticated using SSH keys. Removing the option of password authentication enforces more secure methods while removing the risks inherent with passwords.",
	Impact:      "Using password authentication is less secure that ssh keys may result in compromised servers",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

