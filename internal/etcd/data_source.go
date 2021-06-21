package etcd

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "",
		ReadContext: clusterDataSourceRead,
	}
}

func clusterDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)
	clusters, err := client.Cluster.MemberList(context.Background())
	if err != nil {
		return diag.FromErr(err)
	}
	fmt.Println(clusters)

	return nil
}
