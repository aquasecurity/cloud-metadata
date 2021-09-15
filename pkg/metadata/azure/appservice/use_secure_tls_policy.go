package appservice

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-AZU-0006",
	Title:       "Web App uses latest TLS version",
	Description: "Use a more recent TLS/SSL policy for the App Service",
	Impact:      "The minimum TLS version for apps should be TLS1_2",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

