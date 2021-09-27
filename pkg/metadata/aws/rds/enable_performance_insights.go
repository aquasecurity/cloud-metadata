package rds

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnablePerformanceInsights = metadata.Metadata{
	ID:          "AVD-AWS-0078",
	Title:       "Encryption for RDS Performance Insights should be enabled.",
	Description: "When enabling Performance Insights on an RDS cluster or RDS DB Instance, and encryption key should be provided.

The encryption key specified in `performance_insights_kms_key_id` references a KMS ARN",
	Impact:      "Data can be read from the RDS Performance Insights if it is compromised",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.htm", 
	},
}

