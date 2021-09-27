package elbv2

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var HttpNotUsed = metadata.Metadata{
	ID:          "AVD-AWS-0054",
	Title:       "Use of plain HTTP.",
	Description: "Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection, they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.",
	Impact:      "Your traffic is not protected",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://www.cloudflare.com/en-gb/learning/ssl/why-is-http-not-secure/", 
	},
}

