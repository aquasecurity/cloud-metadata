package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var MetadataEndpointsDisabled = metadata.Metadata{
	ID:          "AVD-GCP-0031",
	Title:       "Legacy metadata endpoints enabled.",
	Description: "The Compute Engine instance metadata server exposes legacy v0.1 and v1beta1 endpoints, which do not enforce metadata query headers. 

This is a feature in the v1 APIs that makes it more difficult for a potential attacker to retrieve instance metadata. 

Unless specifically required, we recommend you disable these legacy APIs.

When setting the metadata block, the default value for disable-legacy-endpoints is set to true, they should not be explicitly enabled.",
	Impact:      "Legacy metadata endpoints don't require metadata headers",
	Severity:    "HIGH",
	Links:       []string {
		"https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#protect_node_metadata_default_for_112", 
	},
}

