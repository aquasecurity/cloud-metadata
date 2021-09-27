package loadbalancing

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforceHttps = metadata.Metadata{
	ID:          "AVD-DIG-0004",
	Title:       "The load balancer forwarding rule is using an insecure protocol as an entrypoint",
	Description: "Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection, they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.",
	Impact:      "Your inbound traffic is not protected",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.digitalocean.com/products/networking/load-balancers/", 
	},
}

