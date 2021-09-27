package autoscaling

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoPublicIp = metadata.Metadata{
	ID:          "AVD-AWS-0009",
	Title:       "A resource has a public IP address.",
	Description: "You should limit the provision of public IP addresses for resources. Resources should not be exposed on the public internet, but should have access limited to consumers required for the function of your application.",
	Impact:      "The instance or configuration is publicly accessible",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-instance-addressing.html", 
	},
}

