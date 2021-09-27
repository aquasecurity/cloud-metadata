package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicEgress = metadata.Metadata{
	ID:          "AVD-DIG-0001",
	Title:       "The firewall has an outbound rule with open access",
	Description: "Opening up ports to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that explicitly require it where possible.",
	Impact:      "The port is exposed for ingress from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.digitalocean.com/products/networking/firewalls/how-to/configure-rules/", 
	},
}

