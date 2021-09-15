package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RequireSymbolsInPasswords = metadata.Metadata{
	ID:          "AVD-AWS-0060",
	Title:       "IAM Password policy should have requirement for at least one symbol in the password.",
	Description: "IAM account password policies should ensure that passwords content including a symbol.",
	Impact:      "Short, simple passwords are easier to compromise",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html#password-policy-details", 
	},
}

