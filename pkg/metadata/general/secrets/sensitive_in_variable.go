package secrets

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SensitiveInVariable = metadata.Metadata{
	ID:          "AVD-GEN-0004",
	Title:       "Potentially sensitive data stored in "default" value of variable.",
	Description: "Sensitive attributes such as passwords and API tokens should not be available in your templates, especially in a plaintext form. You can declare variables to hold the secrets, assuming you can provide values for those variables in a secure fashion. Alternatively, you can store these secrets in a secure secret store, such as AWS KMS.

*NOTE: It is also recommended to store your Terraform state in an encrypted form.*",
	Impact:      "Default values could be exposing sensitive data",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

