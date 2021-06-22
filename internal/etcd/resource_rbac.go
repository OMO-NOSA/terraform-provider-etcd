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
				ForceNew: true,
			},
		},
	}
}

func RoleGrantResource() *schema.Resource {
	return &schema.Resource{
		Description:   "",
		CreateContext: GrantUserRole,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func GrantUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("role_name").(string)
	userName := d.Get("username").(string)
	_, err := client.UserGrantRole(ctx, userName, roleName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(roleName)
	return nil

}

func RevokeUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("role_name").(string)
	userName := d.Get("username").(string)
	_, err := client.UserRevokeRole(ctx, userName, roleName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func RolePermissionResource() *schema.Resource {
	return &schema.Resource{
		Description: "",
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

	resp, err := client.RoleGet(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("permissions", resp.Perm); err != nil {
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
