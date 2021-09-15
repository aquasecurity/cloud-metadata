package compute

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoSerialPort = metadata.Metadata{
	ID:          "AVD-GCP-0016",
	Title:       "Disable serial port connectivity for all instances",
	Description: "When serial port access is enabled, the access is not governed by network security rules meaning the port can be exposed publicly.",
	Impact:      "Unrestricted network access to the serial console of the instance",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

