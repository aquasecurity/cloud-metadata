package functionapp

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AuthenticationEnabled = metadata.Metadata{
	ID:          "AVD-AZU-0027",
	Title:       "Function App authentication is activated",
	Description: "Enabling authentication ensures that all communications in the application are authenticated. The auth_settings block needs to be filled out with the appropriate auth backend settings",
	Impact:      "Anonymous HTTP requests will be accepted",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

