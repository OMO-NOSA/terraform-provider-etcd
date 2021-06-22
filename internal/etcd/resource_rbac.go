package etcd

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clientv3 "go.etcd.io/etcd/client/v3"
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

func RoleResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("name").(string)

	_, err := client.RoleAdd(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	d.Set("name", roleName)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
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

func RoleGrantResource() *schema.Resource {
	return &schema.Resource{
		Description:   "",
		CreateContext: GrantUserRole,
		ReadContext:   Users,
		DeleteContext: RevokeUserRole,
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
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil

}

func Users(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)
	users, err := client.UserList(ctx)

	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("users", users.Users); err != nil {
		return diag.FromErr(err)
	}
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
		CreateContext: GrantRolePermission,
		DeleteContext: RevokeRolePermission,
		ReadContext:   NotImplemented,
		Description:   "",
		Schema: map[string]*schema.Schema{
			"role_name": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"key": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"range": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"permission": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func GrantRolePermission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)
	rangeEnd := d.Get("range").(string)
	roleName := d.Get("role_name").(string)
	key := d.Get("key").(string)
	permission := d.Get("permission").(string)

	perm, err := clientv3.StrToPermissionType(permission)
	if err != nil {
		return diag.FromErr(err)
	}

	rangePrefix := clientv3.GetPrefixRangeEnd(rangeEnd)

	_, err = client.RoleGrantPermission(ctx, roleName, key, rangePrefix, perm)

	if err != nil {
		diag.FromErr(err)
	}
	return nil
}

func RevokeRolePermission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)
	rangeEnd := d.Get("range").(string)
	roleName := d.Get("role_name").(string)
	key := d.Get("key").(string)

	_, err := client.RoleRevokePermission(ctx, roleName, key, rangeEnd)

	if err != nil {
		diag.FromErr(err)
	}
	return nil

}

func NotImplemented(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil

}
