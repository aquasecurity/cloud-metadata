package rds

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var BackupRetentionSpecified = metadata.Metadata{
	ID:          "AVD-AWS-0077",
	Title:       "RDS Cluster and RDS instance should have backup retention longer than default 1 day",
	Description: "RDS backup retention for clusters defaults to 1 day, this may not be enough to identify and respond to an issue. Backup retention periods should be set to a period that is a balance on cost and limiting risk.",
	Impact:      "Potential loss of data and short opportunity for recovery",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupRetention", 
	},
}

