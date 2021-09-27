package documentdb

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableLogExport = metadata.Metadata{
	ID:          "AVD-AWS-0020",
	Title:       "DocumentDB logs export should be enabled",
	Description: "Document DB does not have auditing by default. To ensure that you are able to accurately audit the usage of your DocumentDB cluster you should enable export logs.",
	Impact:      "Limited visibility of audit trail for changes to the DocumentDB",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

