package lambda

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var RestrictSourceArn = metadata.Metadata{
	ID:          "AVD-AWS-0067",
	Title:       "Ensure that lambda function permission has a source arn specified",
	Description: "When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to. 

Without this, any resource from principal will be granted permission – even if that resource is from another account. 

For S3, this should be the ARN of the S3 Bucket. For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule. For API Gateway, this should be the ARN of the API",
	Impact:      "Not providing the source ARN allows any resource from principal, even from other accounts",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html", 
	},
}

