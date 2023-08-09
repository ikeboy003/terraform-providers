package test

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func MockProvider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"custom_vm_domain_join": resourceVMJoinDomain(),
		},
	}
}
