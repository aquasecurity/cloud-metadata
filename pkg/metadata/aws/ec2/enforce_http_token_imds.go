package ec2

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnforceHttpTokenImds = metadata.Metadata{
	ID:          "AVD-AWS-0028",
	Title:       "aws_instance should activate session tokens for Instance Metadata Service.",
	Description: "IMDS v2 (Instance Metadata Service) introduced session authentication tokens which improve security when talking to IMDS.
By default aws_instance resource sets IMDS session auth tokens to be optional. 
To fully protect IMDS you need to enable session tokens by using metadata_options block and its http_tokens variable set to required.",
	Impact:      "Instance metadata service can be interacted with freely",
	Severity:    "HIGH",
	Links:       []string {
		"https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service", 
	},
}

