package neptune

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableLogExport = metadata.Metadata{
	ID:          "AVD-AWS-0075",
	Title:       "Nepture logs export should be enabled",
	Description: "Neptune does not have auditing by default. To ensure that you are able to accurately audit the usage of your Neptune instance you should enable export logs.",
	Impact:      "Limited visibility of audit trail for changes to Neptune",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

