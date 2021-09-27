package vpc

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoDefaultVpc = metadata.Metadata{
	ID:          "AVD-AWS-0101",
	Title:       "AWS best practice to not use the default VPC for workflows",
	Description: "Default VPC does not have a lot of the critical security features that standard VPC comes with, new resources should not be created in the default VPC and it should not be present in the Terraform.",
	Impact:      "The default VPC does not have critical security features applied",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html", 
	},
}

