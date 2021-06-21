package etcd

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "",

		CreateContext: AuthResourceCreateUser,
		ReadContext:   AuthResourceGetUser,
		UpdateContext: AuthResourceUpdateUser,
		DeleteContext: AuthResourceDeleteUser,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func AuthResourceCreateUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)
	passWord := d.Get("password").(string)

	if passWord == "" {
		errmsg := errors.New("password cannot be empty")
		return diag.FromErr(errmsg)
	}

	_, err := client.UserAdd(ctx, userName, passWord)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId(userName)
	return nil
}

func AuthResourceDeleteUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)

	_, err := client.UserDelete(ctx, userName)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func AuthResourceUpdateUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)
	passWord := d.Get("password").(string)

	if passWord == "" {
		errmsg := errors.New("password cannot be empty")
		return diag.FromErr(errmsg)
	}

	_, err := client.UserChangePassword(ctx, userName, passWord)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func AuthResourceGetUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)

	_, err := client.UserGet(ctx, userName)
	if err != nil {
		diag.FromErr(err)
	}

	return nil
}
