package eks

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableControlPlaneLogging = metadata.Metadata{
	ID:          "AVD-AWS-0038",
	Title:       "EKS Clusters should have cluster control plane logging turned on",
	Description: "By default cluster control plane logging is not turned on. Logging is available for audit, api, authenticator, controllerManager and scheduler. All logging should be turned on for cluster control plane.",
	Impact:      "Logging provides valuable information about access and usage",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html", 
	},
}

