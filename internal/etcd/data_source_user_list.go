package etcd

import (
	"context"
	"strconv"
	"time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UsersDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "",
		ReadContext: usersDataSourceRead,
		Schema: map[string]*schema.Schema{
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func usersDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)


	users, err := client.UserList(ctx)

	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("users", users.Users); err != nil {
		
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
