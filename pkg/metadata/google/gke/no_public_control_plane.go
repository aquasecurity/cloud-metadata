package gke

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicControlPlane = metadata.Metadata{
	ID:          "AVD-GCP-0034",
	Title:       "GKE Control Plane should not be publicly accessible",
	Description: "The GKE control plane is exposed to the public internet by default.",
	Impact:      "GKE control plane exposed to public internet",
	Severity:    "HIGH",
	Links:       []string {
		
	},
}

