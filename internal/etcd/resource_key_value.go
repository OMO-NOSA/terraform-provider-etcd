package etcd

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/clientv3util"
)

func KvResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "",

		CreateContext: KvResourceCreate,
		ReadContext:   KvResourceRead,
		//UpdateContext: resourceUpdate,
		DeleteContext: KvResourceDelete,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func KvResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*apiClient).Client

	key := d.Get("key").(string)
	value := d.Get("value").(string)

	kvc := clientv3.NewKV(client)

	_, err := kvc.Txn(ctx).If(clientv3util.KeyMissing(key)).Then(clientv3.OpPut(key, value)).Commit()

	if err != nil {
		switch err {
		case context.Canceled:
			errmsg := fmt.Errorf("ctx is canceled by another routine: %v", err)
			return diag.FromErr(errmsg)
		case context.DeadlineExceeded:
			errmsg := fmt.Errorf("ctx is attached with a deadline is exceeded: %v", err)
			return diag.FromErr(errmsg)
		case rpctypes.ErrEmptyKey:
			errmsg := fmt.Errorf("client-side error: %v", err)
			return diag.FromErr(errmsg)
		default:
			errmsg := fmt.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
			return diag.FromErr(errmsg)
		}

	}
	d.SetId(key)

	return diags
}

func KvResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	id := d.Id()

	response, err := client.Get(ctx, id)
	if err != nil {
		return diag.FromErr(err)

	}
	d.Set("key", response.Kvs)
	return nil
}

func KvResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).Client

	id := d.Id()

	kvc := clientv3.NewKV(client)

	_, err := kvc.Txn(ctx).
		If(clientv3util.KeyExists(id)).
		Then(clientv3.OpDelete(id)).
		Commit()

	if err != nil {
		return diag.FromErr(err)

	}
	d.SetId("")
	return nil
}
