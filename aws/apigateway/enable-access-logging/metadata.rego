package cloud-metadata.aws.apigateway.enable-access-logging

__rego_metadata__ := {
	"id": "AWS061",
	"apiVersion": 2,
	"version": 1,
	"title": "API Gateway stages for V1 and V2 should have access logging enabled",
	"description": "Ensures that Amazon API Gateway API stages have Amazon CloudWatch Logs enabled.",
	"custom": {
		"severity": "MEDIUM",
		"possibleImpact": "Logging provides vital information about access and usage.",
		"urls": ["https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html"],
		"apis": ["APIGateway:getRestApis", "APIGateway:getStages"],
		"iac": true,
		"cspm": true
	}
}