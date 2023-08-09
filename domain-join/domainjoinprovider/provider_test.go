package domainjoinprovider_test

import (
	DomainJoinProvider "domain-join/domainjoinprovider"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestResourceVMJoinDomainCreate(t *testing.T) {

	// Create a ResourceData instance with test data
	data := schema.TestResourceDataRaw(t, DomainJoinProvider.Resource().Schema, nil)
	// Interface is used to pas
	m := interface{}(nil)
	data.Set("vm_name", "test-vm")
	data.Set("domain_name", "example.com")
	data.Set("ou", "OU=Computers,DC=example,DC=com")
	data.Set("username", "testuser")
	data.Set("password", "testpassword")

	err := DomainJoinProvider.ResourceVMJoinDomainCreate(data, m)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Add assertions to verify the behavior of your function
	// For example, you could check if the ResourceData ID was set correctly after successful execution.
	if id := data.Id(); id != "test-vm" {
		t.Errorf("expected resource ID to be 'test-vm', got '%s'", id)
	}
}

func TestResourceVMJoinDomainDelete(t *testing.T) {

	// Create a ResourceData instance with test data
	data := schema.TestResourceDataRaw(t, DomainJoinProvider.Resource().Schema, nil)
	data.SetId("test-vm")

	// Call the function being tested
	err := DomainJoinProvider.ResourceVMJoinDomainDelete(data, nil)

	// Assert the result
	if err != nil {
		t.Errorf("Error unjoining VM from domain: %v", err)
	}
}
