package etcd


import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "",

		CreateContext: AuthenticateUserResource,
		ReadContext:   AuthenticateUserReadResource,
		UpdateContext: nil,
		DeleteContext: AuthenticateUserDeleteResource,

		Schema: map[string]*schema.Schema{
			"auth_status": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_auth_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
		},
		},
	}
}

func AuthenticateUserResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	isAuthEnabled := d.Get("is_auth_enabled").(bool)

	if !isAuthEnabled {
		client.AuthDisable(ctx)
		d.Set("auth_status", false)
		status, err := client.AuthStatus(ctx)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("auth_status", status.Enabled)
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		return nil
	}

	_, err := client.AuthEnable(ctx)

	if err != nil {
		return diag.FromErr(err)
	}

	status, err := client.AuthStatus(ctx)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("auth_status", status.Enabled)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func AuthenticateUserReadResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	status :=  d.Get("auth_status")
	d.Set("auth_status", status)
	return nil 
}

func AuthenticateUserDeleteResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil 
}
