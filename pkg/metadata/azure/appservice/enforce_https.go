package appservice

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforceHttps = metadata.Metadata{
	ID:          "AVD-AZU-0004",
	Title:       "Ensure the Function App can only be accessed via HTTPS. The default is false.",
	Description: "By default, clients can connect to function endpoints by using both HTTP or HTTPS. You should redirect HTTP to HTTPs because HTTPS uses the SSL/TLS protocol to provide a secure connection, which is both encrypted and authenticated.",
	Impact:      "Anyone can access the Function App using HTTP.",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/app-service/configure-ssl-bindings#enforce-https", "https://docs.microsoft.com/en-us/azure/azure-functions/security-concepts", 
	},
}

