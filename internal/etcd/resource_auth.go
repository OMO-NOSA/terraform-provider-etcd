package etcd

import (
	"context"
	"errors"
	"strconv"
	"time"

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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func AuthResourceCreateUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)
	passWord := d.Get("password").(string)

	if (passWord == "" || len(passWord) < 9) {
		errmsg := errors.New("Validate Password Strength")
		return diag.FromErr(errmsg)
	}

	_, err := client.UserAdd(ctx, userName, passWord)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("username", userName)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func AuthResourceDeleteUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)

	_, err := client.UserDelete(ctx, userName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func AuthResourceUpdateUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)
	passWord := d.Get("password").(string)

	if (passWord == "" || len(passWord) < 9) {
		errmsg := errors.New("Validate Password Strength")
		return diag.FromErr(errmsg)
	}

	_, err := client.UserChangePassword(ctx, userName, passWord)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("last_updated", time.Now().Format(time.RFC850))
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func AuthResourceGetUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	userName := d.Get("username").(string)

	resp, err := client.UserGet(ctx, userName)

	if err != nil {
		return diag.FromErr(err)
	}
	roles := []string{}

	for _, role := range resp.Roles {
		roles = append(roles, role)
	}
	if err := d.Set("roles", roles); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
