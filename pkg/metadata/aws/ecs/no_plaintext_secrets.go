package ecs

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPlaintextSecrets = metadata.Metadata{
	ID:          "AVD-AWS-0036",
	Title:       "Task definition defines sensitive environment variable(s).",
	Description: "You should not make secrets available to a user in plaintext in any scenario. Secrets can instead be pulled from a secure secret storage system by the service requiring them.",
	Impact:      "Sensitive data could be exposed in the AWS Management Console",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/systems-manager/latest/userguide/integration-ps-secretsmanager.html", "https://www.vaultproject.io/", 
	},
}

