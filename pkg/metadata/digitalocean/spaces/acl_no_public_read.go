package spaces

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AclNoPublicRead = metadata.Metadata{
	ID:          "AVD-DIG-0005",
	Title:       "Spaces bucket or bucket object has public read acl set",
	Description: "Space bucket and bucket object permissions should be set to deny public access unless explicitly required.",
	Impact:      "The contents of the space can be accessed publicly",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.digitalocean.com/reference/api/spaces-api/#access-control-lists-acls", 
	},
}

