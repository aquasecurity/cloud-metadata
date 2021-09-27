package container

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var ConfiguredNetworkPolicy = metadata.Metadata{
	ID:          "AVD-AZU-0012",
	Title:       "Ensure AKS cluster has Network Policy configured",
	Description: "The Kubernetes object type NetworkPolicy should be defined to have opportunity allow or block traffic to pods, as in a Kubernetes cluster configured with default settings, all pods can discover and communicate with each other without any restrictions.",
	Impact:      "No network policy is protecting the AKS cluster",
	Severity:    "HIGH",
	Links:       []string {
		"https://kubernetes.io/docs/concepts/services-networking/network-policies", 
	},
}

