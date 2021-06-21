package etcd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RoleResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "",

		CreateContext: RoleResourceCreate,
		ReadContext:   RoleResourceRead,
		//UpdateContext: resourceUpdate,
		DeleteContext: RoleResourceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func RolePermissionResource() *schema.Resource {
	return &schema.Resource{
		Description: "",
		CreateContext: ,
		
	}
}

func RoleResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("name").(string)

	_, err := client.RoleAdd(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId(roleName)
	return nil
}

func RoleResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("name").(string)

	_, err := client.RoleGet(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	return nil
}

func RoleResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("name").(string)

	_, err := client.RoleDelete(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
