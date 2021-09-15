package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableInTransitEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0043",
	Title:       "Elasticsearch domain uses plaintext traffic for node to node communication.",
	Description: "Traffic flowing between Elasticsearch nodes should be encrypted to ensure sensitive data is kept private.",
	Impact:      "In transit data between nodes could be read if intercepted",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/ntn.html", 
	},
}

