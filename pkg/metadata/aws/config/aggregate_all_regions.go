package config

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var AggregateAllRegions = metadata.Metadata{
	ID:          "AVD-AWS-0019",
	Title:       "Config configuration aggregator should be using all regions for source",
	Description: "The configuration aggregator should be configured with all_regions for the source. 

This will help limit the risk of any unmonitored configuration in regions that are thought to be unused.",
	Impact:      "Sources that aren't covered by the aggregator are not include in the configuration",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/config/latest/developerguide/aggregate-data.html", 
	},
}

