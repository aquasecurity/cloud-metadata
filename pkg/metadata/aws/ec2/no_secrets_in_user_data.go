package ec2

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var NoSecretsInUserData = metadata.Metadata{
	ID:          "AVD-AWS-0029",
	Title:       "User data for EC2 instances must not contain sensitive AWS keys",
	Description: "EC2 instance data is used to pass start up information into the EC2 instance. This userdata must not contain access key credentials. Instead use an IAM Instance Profile assigned to the instance to grant access to other AWS Services.",
	Impact:      "User data is visible through the AWS Management console",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-add-user-data.html", 
	},
}

