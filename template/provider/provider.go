package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// give a name to the provider @resource_name
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"resource_name": resouce(),
		},
	}
}

/*
To add more more parameters, add below after schema with a trailing comma
ex.
	"param1": {
					Type:     schema.TypeInt,
					Required: true,
				},

Using this example
	resource "resource_name" "example" {
	  param1        = 1
	}

*/

// Rename Function to Match your needs
func resouce() *schema.Resource {
	return &schema.Resource{
		Create: createResource,
		Read:   readResource,
		Update: updateResource,
		Delete: deleteResource,
		Exists: resourceExists,
		Schema: map[string]*schema.Schema{
			"param1": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func createResource(d *schema.ResourceData, m interface{}) error {
	//Implement logic to create resource

	return nil
}

func resourceExists(d *schema.ResourceData, m interface{}) (bool, error) {
	//Implement Does this resource exists?
	return true, nil
}

func deleteResource(*schema.ResourceData, interface{}) error {
	//Implement logic to delete a resource
	return nil
}

func updateResource(*schema.ResourceData, interface{}) error {
	//Implement logic to update a resource
	return nil
}

func readResource(*schema.ResourceData, interface{}) error {
	//Implement logic to read a resource
	return nil
}
