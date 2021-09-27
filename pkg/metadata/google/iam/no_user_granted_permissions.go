package iam

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoUserGrantedPermissions = metadata.Metadata{
	ID:          "AVD-GCP-0048",
	Title:       "IAM granted directly to user.",
	Description: "Permissions should not be directly granted to users, you identify roles that contain the appropriate permissions, and then grant those roles to the user. 

Granting permissions to users quickly become unwieldy and complex to make large scale changes to remove access to a particular resource.

Permissions should be granted on roles, groups, services accounts instead.",
	Impact:      "Users shouldn't have permissions granted to them directly",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/iam/docs/overview#permissions", "https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy", 
	},
}

