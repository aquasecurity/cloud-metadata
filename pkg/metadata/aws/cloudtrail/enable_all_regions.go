package cloudtrail

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableAllRegions = metadata.Metadata{
	ID:          "AVD-AWS-0014",
	Title:       "Cloudtrail should be enabled in all regions regardless of where your AWS resources are generally homed",
	Description: "When creating Cloudtrail in the AWS Management Console the trail is configured by default to be multi-region, this isn't the case with the Terraform resource. Cloudtrail should cover the full AWS account to ensure you can track changes in regions you are not actively operting in.",
	Impact:      "Activity could be happening in your account in a different region",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/awscloudtrail/latest/userguide/receive-cloudtrail-log-files-from-multiple-regions.html", 
	},
}

