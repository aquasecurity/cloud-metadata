package repositories

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var Private = metadata.Metadata{
	ID:          "AVD-GIT-0001",
	Title:       "Github repository shouldn't be public.",
	Description: "Github repository should be set to be private.

You can do this by either setting private attribute to 'true' or visibility attribute to 'internal' or 'private'.",
	Impact:      "Anyone can read the contents of the GitHub repository and leak IP",
	Severity:    "CRITICAL",
	Links:       []string {
		"https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/about-repository-visibility", "https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/about-repository-visibility#about-internal-repositories", 
	},
}

