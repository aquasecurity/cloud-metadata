package appservice

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AccountIdentityRegistered = metadata.Metadata{
	ID:          "AVD-AZU-0001",
	Title:       "Web App has registration with AD enabled",
	Description: "Registering the identity used by an App with AD allows it to interact with other services without using username and password",
	Impact:      "Interaction between services can't easily be achieved without username/password",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

