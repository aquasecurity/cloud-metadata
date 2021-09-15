package sql

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var EnablePgTempFileLogging = metadata.Metadata{
	ID:          "AVD-GCP-0052",
	Title:       "Temporary file logging should be enabled for all temporary files.",
	Description: "Temporary files are not logged by default. To log all temporary files, a value of `0` should set in the `log_temp_files` flag - as all files greater in size than the number of bytes set in this flag will be logged.",
	Impact:      "Use of temporary files will not be logged",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://postgresqlco.nf/doc/en/param/log_temp_files/", 
	},
}

