package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPolicyWildcards = metadata.Metadata{
	ID:          "AVD-AWS-0057",
	Title:       "IAM policy should avoid use of wildcards and instead apply the principle of least privilege",
	Description: "You should use the principle of least privilege when defining your IAM policies. This means you should specify each exact permission required without using wildcards, as this could cause the granting of access to certain undesired actions, resources and principals.",
	Impact:      "Overly permissive policies may grant access to sensitive resources",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html", 
	},
}

