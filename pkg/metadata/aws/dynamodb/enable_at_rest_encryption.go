package dynamodb

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAtRestEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0023",
	Title:       "DAX Cluster should always encrypt data at rest",
	Description: "Amazon DynamoDB Accelerator (DAX) encryption at rest provides an additional layer of data protection by helping secure your data from unauthorized access to the underlying storage.",
	Impact:      "Data can be freely read if compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAXEncryptionAtRest.html", "https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html", 
	},
}

