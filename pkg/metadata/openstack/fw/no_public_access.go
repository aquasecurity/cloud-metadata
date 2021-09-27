package fw

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicAccess = metadata.Metadata{
	ID:          "AVD-OSK-0002",
	Title:       "A firewall rule allows traffic from/to the public internet",
	Description: "Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.",
	Impact:      "Exposure of infrastructure to the public internet",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

