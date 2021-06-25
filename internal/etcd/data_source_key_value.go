package etcd

import (
	"context"
	"errors"
	"time"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KeyValueDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "",
		ReadContext: keyValueDataSourceRead,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func keyValueDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)
	key := d.Get("key").(string)
	if key == "" {
		errmsg := errors.New("key is empty")
		return diag.FromErr(errmsg)

	}
	value, err := client.Get(ctx, key)
	if err != nil {
		return diag.FromErr(err)
	}

	var keyValue string 

	if len(value.Kvs) > 0 {
		keyValue = string(value.Kvs[0].Value)

	} 

	if err := d.Set("value", keyValue); err != nil {
		return diag.FromErr(err)

	}

	if err := d.Set("key", key); err != nil {
		return diag.FromErr(err)

	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
