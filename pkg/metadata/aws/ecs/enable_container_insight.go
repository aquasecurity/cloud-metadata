package ecs

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableContainerInsight = metadata.Metadata{
	ID:          "AVD-AWS-0034",
	Title:       "ECS clusters should have container insights enabled",
	Description: "Cloudwatch Container Insights provide more metrics and logs for container based applications and micro services.",
	Impact:      "Not all metrics and logs may be gathered for containers when Container Insights isn't enabled",
	Severity:    "LOW",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights.html", 
	},
}

