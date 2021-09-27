package workspace

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableDiskEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0109",
	Title:       "Root and user volumes on Workspaces should be encrypted",
	Description: "Workspace volumes for both user and root should be encrypted to protect the data stored on them.",
	Impact:      "Data can be freely read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/workspaces/latest/adminguide/encrypt-workspaces.html", 
	},
}

