package autoscaling

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAtRestEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0008",
	Title:       "Launch configuration with unencrypted block device.",
	Description: "Blocks devices should be encrypted to ensure sensitive data is held securely at rest.",
	Impact:      "The block device could be compromised and read from",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/RootDeviceStorage.html", 
	},
}

