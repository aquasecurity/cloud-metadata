package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIngress = metadata.Metadata{
	ID:          "AVD-DIG-0002",
	Title:       "The firewall has an inbound rule with open access",
	Description: "Opening up ports to connect out to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that are explicitly required where possible.",
	Impact:      "Your port is exposed to the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.digitalocean.com/products/networking/firewalls/how-to/configure-rules/", 
	},
}

