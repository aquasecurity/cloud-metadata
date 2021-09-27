package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SetMaxPasswordAge = metadata.Metadata{
	ID:          "AVD-AWS-0062",
	Title:       "IAM Password policy should have expiry less than or equal to 90 days.",
	Description: "IAM account password policies should have a maximum age specified. 
		
The account password policy should be set to expire passwords after 90 days or less.",
	Impact:      "Long life password increase the likelihood of a password eventually being compromised",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html#password-policy-details", 
	},
}

