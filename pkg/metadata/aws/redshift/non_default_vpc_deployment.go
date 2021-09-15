package redshift

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NonDefaultVpcDeployment = metadata.Metadata{
	ID:          "AVD-AWS-0085",
	Title:       "Redshift cluster should be deployed into a specific VPC",
	Description: "Redshift clusters that are created without subnet details will be created in EC2 classic mode, meaning that they will be outside of a known VPC and running in tennant.

In order to benefit from the additional security features achieved with using an owned VPC, the subnet should be set.",
	Impact:      "Redshift cluster does not benefit from VPC security if it is deployed in EC2 classic mode",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/redshift/latest/mgmt/managing-clusters-vpc.html", 
	},
}

