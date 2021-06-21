package provider

import (
	"os"
	"testing"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.

func TestEctdEndpoints(t *testing.T) {
	if err := os.Getenv("ENDPOINTS"); err == "" {
		t.Fatal("HASHICUPS_USERNAME must be set for acceptance tests")
	}

}
