package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforcePodSecurityPolicy = metadata.Metadata{
	ID:          "AVD-GCP-0030",
	Title:       "Pod security policy enforcement not defined.",
	Description: "By default, Pods in Kubernetes can operate with capabilities beyond what they require. You should constrain the Pod's capabilities to only those required for that workload.

Kubernetes offers controls for restricting your Pods to execute with only explicitly granted capabilities. 

Pod Security Policy allows you to set smart defaults for your Pods, and enforce controls you want to enable across your fleet. 

The policies you define should be specific to the needs of your application",
	Impact:      "Pods could be operating with more permissions than required to be effective",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#admission_controllers", 
	},
}

