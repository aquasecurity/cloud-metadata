package network

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var SshBlockedFromInternet = metadata.Metadata{
	ID:          "AVD-AZU-0043",
	Title:       "SSH access should not be accessible from the Internet, should be blocked on port 22",
	Description: "SSH access can be configured on either the network security group or in the network security group rule. 

SSH access should not be permitted from the internet (*, 0.0.0.0, /0, internet, any)",
	Impact:      "Its dangerous to allow SSH access from the internet",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

