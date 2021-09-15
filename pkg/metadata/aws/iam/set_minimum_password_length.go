package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SetMinimumPasswordLength = metadata.Metadata{
	ID:          "AVD-AWS-0063",
	Title:       "IAM Password policy should have minimum password length of 14 or more characters.",
	Description: "IAM account password policies should ensure that passwords have a minimum length. 

The account password policy should be set to enforce minimum password length of at least 14 characters.",
	Impact:      "Short, simple passwords are easier to compromise",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html#password-policy-details", 
	},
}

