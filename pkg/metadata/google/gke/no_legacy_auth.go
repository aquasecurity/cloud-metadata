package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoLegacyAuth = metadata.Metadata{
	ID:          "AVD-GCP-0032",
	Title:       "Clusters should use client certificates for authentication",
	Description: "Client certificates are the most secure and recommended method of authentication.",
	Impact:      "Less secure authentication method in use",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#restrict_authn_methods", 
	},
}

