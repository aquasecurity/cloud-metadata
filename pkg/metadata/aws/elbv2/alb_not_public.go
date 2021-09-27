package elbv2

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AlbNotPublic = metadata.Metadata{
	ID:          "AVD-AWS-0053",
	Title:       "Load balancer is exposed to the internet.",
	Description: "There are many scenarios in which you would want to expose a load balancer to the wider internet, but this check exists as a warning to prevent accidental exposure of internal assets. You should ensure that this resource should be exposed publicly.",
	Impact:      "The load balancer is exposed on the internet",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

