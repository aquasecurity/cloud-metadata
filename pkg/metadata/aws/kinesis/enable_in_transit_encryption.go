package kinesis

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableInTransitEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0064",
	Title:       "Kinesis stream is unencrypted.",
	Description: "Kinesis streams should be encrypted to ensure sensitive data is kept private. Additionally, non-default KMS keys should be used so granularity of access control can be ensured.",
	Impact:      "Intercepted data can be read in transit",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/streams/latest/dev/server-side-encryption.html", 
	},
}

