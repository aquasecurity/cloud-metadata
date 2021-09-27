package codebuild

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnableEncryption = metadata.Metadata{
	ID:          "AVD-AWS-0018",
	Title:       "CodeBuild Project artifacts encryption should not be disabled",
	Description: "All artifacts produced by your CodeBuild project pipeline should always be encrypted",
	Impact:      "CodeBuild project artifacts are unencrypted",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html", "https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html", 
	},
}

