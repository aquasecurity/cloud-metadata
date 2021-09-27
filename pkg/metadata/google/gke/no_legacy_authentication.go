package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoLegacyAuthentication = metadata.Metadata{
	ID:          "AVD-GCP-0033",
	Title:       "Legacy client authentication methods utilized.",
	Description: "It is recommended to use Service Accounts and OAuth as authentication methods for accessing the master in the container cluster. 

Basic authentication should be disabled by explicitly unsetting the username and password on the master_auth block.",
	Impact:      "Username and password authentication methods are less secure",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#restrict_authn_methods", 
	},
}

