package etcd

import (
	"testing"
)

func TestProvider(test *testing.T) {
	provider := New()
	if err := provider.InternalValidate(); err != nil {
		test.Fatalf("err: %s", err)
	}
}
