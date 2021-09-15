package network

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicEgress = metadata.Metadata{
	ID:          "AVD-K8S-0001",
	Title:       "Public egress should not be allowed via network policies",
	Description: "You should not expose infrastructure to the public internet except where explicitly required",
	Impact:      "Exfiltration of data to the public internet",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

