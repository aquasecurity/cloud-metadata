package mssql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var ThreatAlertEmailSet = metadata.Metadata{
	ID:          "AVD-AZU-0037",
	Title:       "At least one email address is set for threat alerts",
	Description: "SQL Server sends alerts for threat detection via email, if there are no email addresses set then mitigation will be delayed.",
	Impact:      "Nobody will be prompty alerted in the case of a threat being detected",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

