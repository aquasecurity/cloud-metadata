package droplet

import "github.com/aquasecurity/cloud-metadata/pkg/metadata"

var UseSshKeys = metadata.Metadata{
	ID:          "AVD-DIG-0003",
	Title:       "SSH Keys are the preferred way to connect to your droplet, no keys are supplied",
	Description: "When working with a server, you’ll likely spend most of your time in a terminal session connected to your server through SSH. A more secure alternative to password-based logins, SSH keys use encryption to provide a secure way of logging into your server and are recommended for all users.",
	Impact:      "Logging in with username and password is easier to compromise",
	Severity:    "HIGH",
	Links:       []string {
		"https://www.digitalocean.com/community/tutorials/understanding-the-ssh-encryption-and-connection-process", 
	},
}

