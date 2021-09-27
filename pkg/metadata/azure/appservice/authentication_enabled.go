package appservice

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AuthenticationEnabled = metadata.Metadata{
	ID:          "AVD-AZU-0002",
	Title:       "App Service authentication is activated",
	Description: "Enabling authentication ensures that all communications in the application are authenticated. The auth_settings block needs to be filled out with the appropriate auth backend settings",
	Impact:      "Anonymous HTTP requests will be accepted",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

