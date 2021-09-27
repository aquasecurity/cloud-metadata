package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var services = []string {
	"aws",
	"azure",
	"cloudstack",
	"digitalocean",
	"general",
	"github",
	"google",
	"kubernetes",
	"openstack",
	"oracle",
}

func main() {

	tmpl, err := template.New("metadata").Parse(metadataTemplate)
	if err != nil {
		panic(err)
	}

	for _, service := range services {
		paths, err := filepath.Glob(fmt.Sprintf("%s/*/*/metadata.json", service))
		if err != nil {
			panic(err)
		}

		for _, path := range paths {
			generateFile(path, *tmpl)
		}
	}
}

func generateFile(path string, tmpl template.Template) {

	var meta metadata

	pathParts := strings.Split(strings.ReplaceAll(path, "metadata.json", ""), string(os.PathSeparator))

	provider := strings.ReplaceAll(strings.ToTitle(strings.ReplaceAll(pathParts[0], "-", " ")), " ", "")
	service := strings.ReplaceAll(strings.ToTitle(strings.ReplaceAll(pathParts[1], "-", " ")), " ", "")
	rule := strings.ReplaceAll(pathParts[2], "-", "_")

	pkgFilePath := filepath.Join("pkg", "metadata",strings.ToLower(provider), strings.ToLower(service), fmt.Sprintf("%s.go", strings.ToLower(rule)))

	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(content, &meta); err != nil {
		panic(err)
	}

	dir, _ := filepath.Split(pkgFilePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	file, err := os.Create(pkgFilePath)

	if err != nil {
		fmt.Println(dir, pkgFilePath)
		panic(err)
	}

	metaSource := newTemplateSource(service, rule, meta)
	if err := tmpl.Execute(file, metaSource); err != nil {
		panic(err)
	}


}
