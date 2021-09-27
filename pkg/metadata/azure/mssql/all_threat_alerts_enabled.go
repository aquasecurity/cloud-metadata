package mssql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AllThreatAlertsEnabled = metadata.Metadata{
	ID:          "AVD-AZU-0036",
	Title:       "No threat detections are set",
	Description: "SQL Server can alert for security issues including SQL Injection, vulnerabilities, access anomalies and data exfiltration. Ensure none of these are disabled to benefit from the best protection",
	Impact:      "Disabling threat alerts means you are not getting the full benefit of server security protection",
	Severity:    "MEDIUM",
	Links:       []string {
		
	},
}

