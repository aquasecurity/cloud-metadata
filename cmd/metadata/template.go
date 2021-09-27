package main

import (
	"strings"
)

type metadata struct {
	ID          string `json:"id"`
	APIVersion  int    `json:"apiVersion"`
	Version     int    `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Custom      struct {
		Severity       string `json:"severity"`
		PossibleImpact string `json:"possibleImpact"`
		References     []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"references"`
		ExternalToolIds struct {
			Cfsec []string `json:"cfsec"`
			Cspm  []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"cspm"`
			Tfsec []string `json:"tfsec"`
		} `json:"externalToolIds"`
	} `json:"custom"`
}

type metadataTemplateSource struct {
	Service string
	Rule string
	ID string
	Title string
	Description string
	Impact string
	Severity string
	Links []string
}

func newTemplateSource(service, rule string, mdObj metadata) metadataTemplateSource {
	ruleName := strings.ReplaceAll(strings.Title(strings.ReplaceAll(rule, "_", " ")), " ", "")

	var links []string
	for _, reference := range mdObj.Custom.References {
		links = append(links, reference.URL)
	}

	return metadataTemplateSource{
		Service:     strings.ToLower(service),
		Rule:        ruleName,
		ID:          mdObj.ID,
		Title:       mdObj.Title,
		Description: mdObj.Description,
		Impact:      mdObj.Custom.PossibleImpact,
		Severity:    mdObj.Custom.Severity,
		Links:       links,
	}
}

var metadataTemplate = `package {{.Service}}

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var {{.Rule}} = metadata.Metadata{
	ID:          "{{.ID}}",
	Title:       "{{.Title}}",
	Description: "{{.Description}}",
	Impact:      "{{.Impact}}",
	Severity:    "{{.Severity}}",
	Links:       []string {
		{{ range .Links }}"{{.}}", {{ end }}
	},
}

`
