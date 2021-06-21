package etcd

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	
	
func RoleResource() *schema.Resource {
	return &schema.Resource{
		Description:   "",
		CreateContext: RoleResourceCreate,
		DeleteContext: RoleResourceDelete,
		UpdateContext: RoleResourceUpdate,
		Schema: map[string]*schema.Schema{
			"name":   *&schema.Schema{
				"type": TypeString,
			Required: true,
		},
	},
}