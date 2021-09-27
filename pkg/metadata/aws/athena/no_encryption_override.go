package athena

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoEncryptionOverride = metadata.Metadata{
	ID:          "AVD-AWS-0007",
	Title:       "Athena workgroups should enforce configuration to prevent client disabling encryption",
	Description: "Athena workgroup configuration should be enforced to prevent client side changes to disable encryption settings.",
	Impact:      "Clients can ignore encryption requirements",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings.html", 
	},
}

