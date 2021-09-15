package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableIpAliasing = metadata.Metadata{
	ID:          "AVD-GCP-0024",
	Title:       "Clusters should have IP aliasing enabled",
	Description: "IP aliasing allows the reuse of public IPs internally, removing the need for a NAT gateway.",
	Impact:      "Nodes need a NAT gateway to access local services",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

