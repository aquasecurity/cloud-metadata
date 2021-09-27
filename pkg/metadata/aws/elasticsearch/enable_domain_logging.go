package elasticsearch

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableDomainLogging = metadata.Metadata{
	ID:          "AVD-AWS-0042",
	Title:       "Domain logging should be enabled for Elastic Search domains",
	Description: "Amazon ES exposes four Elasticsearch logs through Amazon CloudWatch Logs: error logs, search slow logs, index slow logs, and audit logs. 

Search slow logs, index slow logs, and error logs are useful for troubleshooting performance and stability issues. 

Audit logs track user activity for compliance purposes. 

All the logs are disabled by default.",
	Impact:      "Logging provides vital information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createdomain-configure-slow-logs.html", 
	},
}

