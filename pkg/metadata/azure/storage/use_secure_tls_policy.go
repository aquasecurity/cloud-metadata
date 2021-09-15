package storage

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-AZU-0052",
	Title:       "The minimum TLS version for Storage Accounts should be TLS1_2",
	Description: "Azure Storage currently supports three versions of the TLS protocol: 1.0, 1.1, and 1.2. 

Azure Storage uses TLS 1.2 on public HTTPS endpoints, but TLS 1.0 and TLS 1.1 are still supported for backward compatibility.

This check will warn if the minimum TLS is not set to TLS1_2.",
	Impact:      "The TLS version being outdated and has known vulnerabilities",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/storage/common/transport-layer-security-configure-minimum-version", 
	},
}

