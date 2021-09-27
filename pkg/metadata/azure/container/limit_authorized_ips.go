package container

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var LimitAuthorizedIps = metadata.Metadata{
	ID:          "AVD-AZU-0013",
	Title:       "Ensure AKS has an API Server Authorized IP Ranges enabled",
	Description: "The API server is the central way to interact with and manage a cluster. To improve cluster security and minimize attacks, the API server should only be accessible from a limited set of IP address ranges.",
	Impact:      "Any IP can interact with the API server",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/aks/api-server-authorized-ip-ranges", 
	},
}

