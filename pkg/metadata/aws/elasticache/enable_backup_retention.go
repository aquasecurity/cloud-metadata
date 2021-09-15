package elasticache

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableBackupRetention = metadata.Metadata{
	ID:          "AVD-AWS-0050",
	Title:       "Redis cluster should have backup retention turned on",
	Description: "Redis clusters should have a snapshot retention time to ensure that they are backed up and can be restored if required.",
	Impact:      "Without backups of the redis cluster recovery is made difficult",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-automatic.html", 
	},
}

