package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RequireUppercaseInPasswords = metadata.Metadata{
	ID:          "AVD-AWS-0061",
	Title:       "IAM Password policy should have requirement for at least one uppercase character.",
	Description: ",
IAM account password policies should ensure that passwords content including at least one uppercase character.",
	Impact:      "Short, simple passwords are easier to compromise",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html#password-policy-details", 
	},
}

