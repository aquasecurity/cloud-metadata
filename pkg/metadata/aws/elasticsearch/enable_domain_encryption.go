package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableDomainEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0048",
	Title:       "Elasticsearch domain isn't encrypted at rest.",
	Description: "You should ensure your Elasticsearch data is encrypted at rest to help prevent sensitive information from being read by unauthorised users.",
	Impact:      "Data will be readable if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/encryption-at-rest.html", 
	},
}

