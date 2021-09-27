package dynamodb

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableRecovery = metadata.Metadata{
	ID:          "AVD-AWS-0024",
	Title:       "Point in time recovery should be enabled to protect DynamoDB table",
	Description: "DynamoDB tables should be protected against accidentally or malicious write/delete actions by ensuring that there is adequate protection.

By enabling point-in-time-recovery you can restore to a known point in the event of loss of data.",
	Impact:      "Accidental or malicious writes and deletes can't be rolled back",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/PointInTimeRecovery.html", 
	},
}

