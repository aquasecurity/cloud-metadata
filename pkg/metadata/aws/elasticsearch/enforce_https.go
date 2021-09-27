package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforceHttps = metadata.Metadata{
	ID:          "AVD-AWS-0046",
	Title:       "Elasticsearch doesn't enforce HTTPS traffic.",
	Description: "Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection, they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.",
	Impact:      "HTTP traffic can be intercepted and the contents read",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-data-protection.html", 
	},
}

