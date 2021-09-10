# Cloud Metadata

## What is this

This project contains the metadata for rules that will feed into IaC tools and potentially CSPM.

Rules are broken down into provider/category/rule and will eventually contain markdown files for remediation.

## Things you need to know now

### Schema Validation

A [GitHub action](.github/workflows/validate_metadata.yaml) will run on each PR to ensure that changes to the `metadata.json` files still meet the schema.

### Metadata json

Metadata files define the common information about rules; description, impact, severity and links to the external tools ids (tfsec, cspm, cfsec)

An example metadata.json

```json
{
  "id": "AVD-AWS-0028",
  "apiVersion": 2,
  "version": 1,
  "title": "aws_instance should activate session tokens for Instance Metadata Service.",
  "description": "IMDS v2 (Instance Metadata Service) introduced session authentication tokens which improve security when talking to IMDS.\nBy default eaws_instance resource sets IMDS session auth tokens to be optional. \nTo fully protect IMDS you need to enable session tokens by using metadata_options block and its http_tokens variable set to required.",
  "custom": {
    "severity": "HIGH",
    "possibleImpact": "Instance metadata service can be interacted with freely",
    "references": [
      {
        "title": "title",
        "url": "https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service"
      }
    ],
    "externalToolIds": {
      "cfsec": [
        "aws-ec2-enforce-http-token-imds"
      ],
      "cspm": [
        {
          "id": 692,
          "name": "ec2/ec2MetadataOptions"
        }
      ],
      "tfsec": [
        "aws-ec2-enforce-http-token-imds"
      ]
    }
  }
}
```

The `metadata.json` file for each rule should match the following schema

Type: `object`

<i id="#">path: #</i>

&#36;schema: [http://json-schema.org/draft-04/schema#](http://json-schema.org/draft-04/schema#)

**_Properties_**

 - <b id="#/properties/id">id</b> `required`
	 - Type: `string`
	 - <i id="#/properties/id">path: #/properties/id</i>
 - <b id="#/properties/apiVersion">apiVersion</b> `required`
	 - Type: `integer`
	 - <i id="#/properties/apiVersion">path: #/properties/apiVersion</i>
 - <b id="#/properties/version">version</b> `required`
	 - Type: `integer`
	 - <i id="#/properties/version">path: #/properties/version</i>
 - <b id="#/properties/title">title</b> `required`
	 - Type: `string`
	 - <i id="#/properties/title">path: #/properties/title</i>
 - <b id="#/properties/description">description</b> `required`
	 - Type: `string`
	 - <i id="#/properties/description">path: #/properties/description</i>
 - <b id="#/properties/description_file_path">description_file_path</b>
	 - Type: `string`
	 - <i id="#/properties/description_file_path">path: #/properties/description_file_path</i>
 - <b id="#/properties/custom">custom</b> `required`
	 - Type: `object`
	 - <i id="#/properties/custom">path: #/properties/custom</i>
	 - **_Properties_**
		 - <b id="#/properties/custom/properties/severity">severity</b> `required`
			 - Type: `string`
			 - <i id="#/properties/custom/properties/severity">path: #/properties/custom/properties/severity</i>
		 - <b id="#/properties/custom/properties/possibleImpact">possibleImpact</b> `required`
			 - Type: `string`
			 - <i id="#/properties/custom/properties/possibleImpact">path: #/properties/custom/properties/possibleImpact</i>
		 - <b id="#/properties/custom/properties/remediations">remediations</b>
			 - Type: `array`
			 - <i id="#/properties/custom/properties/remediations">path: #/properties/custom/properties/remediations</i>
				 - **_Items_**
				 - Type: `object`
				 - <i id="#/properties/custom/properties/remediations/items">path: #/properties/custom/properties/remediations/items</i>
				 - **_Properties_**
					 - <b id="#/properties/custom/properties/remediations/items/properties/title">title</b> `required`
						 - Type: `string`
						 - <i id="#/properties/custom/properties/remediations/items/properties/title">path: #/properties/custom/properties/remediations/items/properties/title</i>
					 - <b id="#/properties/custom/properties/remediations/items/properties/remediation_path">remediation_path</b> `required`
						 - Type: `string`
						 - <i id="#/properties/custom/properties/remediations/items/properties/remediation_path">path: #/properties/custom/properties/remediations/items/properties/remediation_path</i>
					 - <b id="#/properties/custom/properties/remediations/items/properties/remediation_type">remediation_type</b> `required`
						 - Type: `string`
						 - <i id="#/properties/custom/properties/remediations/items/properties/remediation_type">path: #/properties/custom/properties/remediations/items/properties/remediation_type</i>
			 - 
		 - <b id="#/properties/custom/properties/references">references</b> `required`
			 - Type: `array`
			 - <i id="#/properties/custom/properties/references">path: #/properties/custom/properties/references</i>
				 - **_Items_**
				 - Type: `object`
				 - <i id="#/properties/custom/properties/references/items">path: #/properties/custom/properties/references/items</i>
				 - **_Properties_**
					 - <b id="#/properties/custom/properties/references/items/properties/title">title</b> `required`
						 - Type: `string`
						 - <i id="#/properties/custom/properties/references/items/properties/title">path: #/properties/custom/properties/references/items/properties/title</i>
					 - <b id="#/properties/custom/properties/references/items/properties/url">url</b> `required`
						 - Type: `string`
						 - <i id="#/properties/custom/properties/references/items/properties/url">path: #/properties/custom/properties/references/items/properties/url</i>
			 - 
		 - <b id="#/properties/custom/properties/refs">refs</b> `required`
			 - Type: `object`
			 - <i id="#/properties/custom/properties/refs">path: #/properties/custom/properties/refs</i>
			 - **_Properties_**
				 - <b id="#/properties/custom/properties/refs/properties/cfsec">cfsec</b>
					 - Type: `array`
					 - <i id="#/properties/custom/properties/refs/properties/cfsec">path: #/properties/custom/properties/refs/properties/cfsec</i>
						 - **_Items_**
						 - Type: `string`
						 - <i id="#/properties/custom/properties/refs/properties/cfsec/items">path: #/properties/custom/properties/refs/properties/cfsec/items</i>
					 - 
				 - <b id="#/properties/custom/properties/refs/properties/cspm">cspm</b>
					 - Type: `array`
					 - <i id="#/properties/custom/properties/refs/properties/cspm">path: #/properties/custom/properties/refs/properties/cspm</i>
						 - **_Items_**
						 - Type: `object`
						 - <i id="#/properties/custom/properties/refs/properties/cspm/items">path: #/properties/custom/properties/refs/properties/cspm/items</i>
						 - **_Properties_**
							 - <b id="#/properties/custom/properties/refs/properties/cspm/items/properties/id">id</b> `required`
								 - Type: `integer`
								 - <i id="#/properties/custom/properties/refs/properties/cspm/items/properties/id">path: #/properties/custom/properties/refs/properties/cspm/items/properties/id</i>
							 - <b id="#/properties/custom/properties/refs/properties/cspm/items/properties/name">name</b> `required`
								 - Type: `string`
								 - <i id="#/properties/custom/properties/refs/properties/cspm/items/properties/name">path: #/properties/custom/properties/refs/properties/cspm/items/properties/name</i>
					 - 
				 - <b id="#/properties/custom/properties/refs/properties/tfsec">tfsec</b>
					 - Type: `array`
					 - <i id="#/properties/custom/properties/refs/properties/tfsec">path: #/properties/custom/properties/refs/properties/tfsec</i>
						 - **_Items_**
						 - Type: `string`
						 - <i id="#/properties/custom/properties/refs/properties/tfsec/items">path: #/properties/custom/properties/refs/properties/tfsec/items</i>
