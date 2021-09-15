package storage

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AllowMicrosoftServiceBypass = metadata.Metadata{
	ID:          "AVD-AZU-0047",
	Title:       "Trusted Microsoft Services should have bypass access to Storage accounts",
	Description: "Some Microsoft services that interact with storage accounts operate from networks that can't be granted access through network rules. 

To help this type of service work as intended, allow the set of trusted Microsoft services to bypass the network rules",
	Impact:      "Trusted Microsoft Services won't be able to access storage account unless rules set to allow",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/storage/common/storage-network-security#trusted-microsoft-services", 
	},
}

