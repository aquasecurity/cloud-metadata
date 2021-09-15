package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSecureTlsPolicy = metadata.Metadata{
	ID:          "AVD-AWS-0108",
	Title:       "An outdated SSL policy is in use by a load balancer.",
	Description: "You should not use outdated/insecure TLS versions for encryption. You should be using TLS v1.2+.",
	Impact:      "The SSL policy is outdated and has known vulnerabilities",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

