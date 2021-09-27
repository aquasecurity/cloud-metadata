package spaces

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var VersioningEnabled = metadata.Metadata{
	ID:          "AVD-DIG-0007",
	Title:       "Spaces buckets should have versioning enabled",
	Description: "Versioning is a means of keeping multiple variants of an object in the same bucket. You can use the Spaces (S3) Versioning feature to preserve, retrieve, and restore every version of every object stored in your buckets. With versioning you can recover more easily from both unintended user actions and application failures.",
	Impact:      "Deleted or modified data would not be recoverable",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html", 
	},
}

