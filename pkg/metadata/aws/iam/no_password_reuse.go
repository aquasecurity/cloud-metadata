package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPasswordReuse = metadata.Metadata{
	ID:          "AVD-AWS-0056",
	Title:       "IAM Password policy should prevent password reuse.",
	Description: "IAM account password policies should prevent the reuse of passwords. 

The account password policy should be set to prevent using any of the last five used passwords.",
	Impact:      "Password reuse increase the risk of compromised passwords being abused",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html#password-policy-details", 
	},
}

