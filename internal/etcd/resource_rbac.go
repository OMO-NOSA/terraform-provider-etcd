package etcd

import (
	"context"
	//"strconv"
	"strings"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func RoleResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Handles Role creation and manages roles lifecycle",

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

	roleName = strings.ToLower(roleName)

	_, err := client.RoleAdd(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	d.Set("name", roleName)
	d.SetId(roleName)
	//d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func RoleResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("name").(string)
	roleName = strings.ToLower(roleName)

	_, err := client.RoleGet(ctx, roleName)
	if err != nil {
		diag.FromErr(err)
	}
	//if err := d.Set("permissions", resp.Perm); err != nil {
	//	diag.FromErr(err)
	//}
	return nil
}

func RoleResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("name").(string)

	roleName = strings.ToLower(roleName)

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
		ReadContext:   NotImplemented,
		DeleteContext: RevokeUserRole,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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

	userName = strings.ToLower(userName)
	roleName = strings.ToLower(roleName)

	_, err := client.UserGrantRole(ctx, userName, roleName)
	if err != nil {
		return diag.FromErr(err)
	}
	//d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.SetId(roleName)
	return nil

}

func Users(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)
	users, err := client.UserList(ctx)

	if err != nil {
		return diag.FromErr(err)
	}
	userList := []string{}

	userList = append(userList, users.Users...)

	if err := d.Set("users", userList); err != nil {
		return diag.FromErr(err)
	}
	return nil

}

func RevokeUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	roleName := d.Get("role_name").(string)
	userName := d.Get("username").(string)
	userName = strings.ToLower(userName)
	roleName = strings.ToLower(roleName)

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
				ForceNew: true,
			},
			"key": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"range": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"permission": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
				ForceNew: true,
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

	roleName = strings.ToLower(roleName)

	perm, err := clientv3.StrToPermissionType(permission)
	if err != nil {
		return diag.FromErr(err)
	}

	rangePrefix := clientv3.GetPrefixRangeEnd(rangeEnd)

	_, err = client.RoleGrantPermission(ctx, roleName, key, rangePrefix, perm)

	if err != nil {
		diag.FromErr(err)
	}
	//d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.SetId(roleName)
	d.Set("role_name", roleName)
	d.Set("key", key)
	d.Set("range", rangePrefix)
	d.Set("permission", perm)
	return nil
}

func RevokeRolePermission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)
	rangeEnd := d.Get("range").(string)
	roleName := d.Get("role_name").(string)
	key := d.Get("key").(string)

	roleName = strings.ToLower(roleName)

	_, err := client.RoleRevokePermission(ctx, roleName, key, rangeEnd)

	if err != nil {
		diag.FromErr(err)
	}
	d.SetId("")
	return nil

}

func NotImplemented(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil

}
