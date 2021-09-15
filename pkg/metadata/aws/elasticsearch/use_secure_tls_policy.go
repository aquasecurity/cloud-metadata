package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-AWS-0047",
	Title:       "Elasticsearch domain endpoint is using outdated TLS policy.",
	Description: "You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.",
	Impact:      "Outdated SSL policies increase exposure to known vulnerabilities",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-data-protection.html", 
	},
}

