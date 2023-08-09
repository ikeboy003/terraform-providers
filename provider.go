package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"custom_vm_domain_join": resourceVMJoinDomain(),
		},
	}
}
func resourceVMJoinDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMJoinDomainCreate,
		Delete: resourceVMJoinDomainDelete,
		Schema: map[string]*schema.Schema{
			"vm_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ou": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceVMJoinDomainCreate(d *schema.ResourceData, m interface{}) error {
	// Type Inference to type string
	vmName := d.Get("vm_name").(string)
	domainName := d.Get("domain_name").(string)
	ou := d.Get("ou").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Construct the PowerShell command to join the VM to the domain
	psCommand := fmt.Sprintf(
		`Add-Computer -DomainName %s -OUPath "%s" -Credential (New-Object System.Management.Automation.PSCredential("%s", (ConvertTo-SecureString -String "%s" -AsPlainText -Force))) -Restart`,
		domainName, ou, username, password,
	)

	// Execute the PowerShell command
	cmd := exec.Command("powershell", "-Command", psCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error joining VM to domain: %v\nOutput: %s", err, output)
	}

	log.Printf("[INFO] VM '%s' successfully joined to domain '%s' under OU '%s'", vmName, domainName, ou)

	d.SetId(vmName) // Set a unique ID for the resource

	return nil
}

func resourceVMJoinDomainDelete(d *schema.ResourceData, m interface{}) error {
	vmName := d.Id()

	// Construct the PowerShell command to unjoin the VM from the domain
	psCommand := fmt.Sprintf(
		"Remove-Computer -UnjoinDomainCredential (Get-Credential) -Force -Verbose",
	)

	// Execute the PowerShell command
	cmd := exec.Command("powershell", "-Command", psCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error unjoining VM from domain: %v\nOutput: %s", err, output)
	}

	log.Printf("[INFO] VM '%s' successfully unjoined from domain", vmName)

	d.SetId("") // Clear the ID to mark the resource as deleted

	return nil
}
