package storage

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-AZU-0050",
	Title:       "Storage containers in blob storage mode should not have public access",
	Description: "Storage container public access should be off. It can be configured for blobs only, containers and blobs or off entirely. The default is off, with no public access.

Explicitly overriding publicAccess to anything other than off should be avoided.",
	Impact:      "Data in the storage container could be exposed publicly",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/storage/blobs/anonymous-read-access-configure?tabs=portal#set-the-public-access-level-for-a-container", 
	},
}

