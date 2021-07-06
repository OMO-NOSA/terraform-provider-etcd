package etcd

import (
	"context"
	// "time"
	// "strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "",
		ReadContext: clusterDataSourceRead,
		Schema: map[string]*schema.Schema{
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_urls": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"client_urls": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func clusterDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)
	clusters, err := client.Cluster.MemberList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	memberList := []interface{}{}
	members := map[string]interface{}{}

	for _, member := range clusters.Members {
		members["id"] = member.ID
		members["name"] = member.Name
		members["peer_urls"] = member.PeerURLs
		members["client_urls"] = member.ClientURLs

		memberList = append(memberList, members)
	}

	if err := d.Set("members", memberList); err != nil {
		return diag.FromErr(err)

	}
	//d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.SetId("cluster_data")
	
	return nil
}
