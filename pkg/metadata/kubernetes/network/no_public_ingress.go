package network

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngress = metadata.Metadata{
	ID:          "AVD-K8S-0002",
	Title:       "Public ingress should not be allowed via network policies",
	Description: "You should not expose infrastructure to the public internet except where explicitly required",
	Impact:      "Exposure of infrastructure to the public internet",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

