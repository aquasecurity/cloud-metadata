package mssql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var ThreatAlertEmailToOwner = metadata.Metadata{
	ID:          "AVD-AZU-0038",
	Title:       "Security threat alerts go to subcription owners and co-administrators",
	Description: "Subscription owners should be notified when there are security alerts. By ensuring the administrators of the account have been notified they can quickly assist in any required remediation",
	Impact:      "Administrators and subscription owners may have a delayed response",
	Severity:    "LOW",
	Links:       []string {
		
	},
}

