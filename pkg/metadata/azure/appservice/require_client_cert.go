package appservice

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RequireClientCert = metadata.Metadata{
	ID:          "AVD-AZU-0005",
	Title:       "Web App accepts incoming client certificate",
	Description: "The TLS mutual authentication technique in enterprise environments ensures the authenticity of clients to the server. If incoming client certificates are enabled only an authenticated client with valid certificates can access the app.",
	Impact:      "Mutual TLS is not being used",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

