package ebs

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableVolumeEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0026",
	Title:       "EBS volumes must be encrypted",
	Description: "By enabling encryption on EBS volumes you protect the volume, the disk I/O and any derived snapshots from compromise if intercepted.",
	Impact:      "Unencrypted sensitive data is vulnerable to compromise.",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html", 
	},
}

