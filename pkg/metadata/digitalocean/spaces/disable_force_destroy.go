package spaces

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var DisableForceDestroy = metadata.Metadata{
	ID:          "AVD-DIG-0006",
	Title:       "Force destroy is enabled on Spaces bucket which is dangerous",
	Description: "Enabling force destroy on a Spaces bucket means that the bucket can be deleted without the additional check that it is empty. This risks important data being accidentally deleted by a bucket removal process.",
	Impact:      "Accidental deletion of bucket objects",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

