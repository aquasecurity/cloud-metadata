package misc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoExposingPlaintextCredentials = metadata.Metadata{
	ID:          "AVD-AWS-0069",
	Title:       "AWS provider has access credentials specified.",
	Description: "The AWS provider block should not contain hardcoded credentials. These can be passed in securely as runtime using environment variables.",
	Impact:      "Exposing the credentials in the Terraform provider increases the risk of secret leakage",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html", 
	},
}

