// main package is the entry point of the Terraform plugin for the Domain Join Provider.
package main

import (
	"domain-join/domainjoinprovider" // Import the custom Domain Join Provider package.

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	// Serve initializes and serves the Terraform plugin.
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			// ProviderFunc returns an instance of the custom Domain Join Provider.

			return domainjoinprovider.Provider()
		},
	})
}
