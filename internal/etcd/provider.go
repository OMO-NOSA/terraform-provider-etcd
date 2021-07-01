package etcd

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	etcd "go.etcd.io/etcd/client/v3"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoints": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				//DefaultFunc: schema.EnvDefaultFunc("ENDPOINTS", []string{"localhost:2379"}),
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("username", ""),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("password", ""),
				
			},
		
		},

		ResourcesMap: map[string]*schema.Resource{
			"etcd_key_value":             KvResource(),
			"etcd_role":                  RoleResource(),
			"etcd_user":                  UserResource(),
			"etcd_grant_user_role":       RoleGrantResource(),
			"etcd_grant_role_permission": RolePermissionResource(),
			"etcd_auth": AuthResource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"etcd_cluster":   ClusterDataSource(),
			"etcd_users":     UsersDataSource(),
			"etcd_key_value": KeyValueDataSource(),
		},
	}

	p.ConfigureContextFunc = configure

	return p

}

type apiClient struct {
	*etcd.Client
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var (
		err   error
		cli   *etcd.Client
	)
	urls := []string{}

	endpoints, ok := d.Get("endpoints").([]string)

	if !ok {
		endpoints := d.Get("endpoints").([]interface{})
		for _, value := range endpoints {
			urls = append(urls, value.(string))
		}

	} else {
		urls = append(urls, endpoints...)
	}
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	
	cli, err = etcd.New(etcd.Config{
		Endpoints:   urls,
		DialTimeout: 5 * time.Second,
		RejectOldCluster: false,
		Username : username,
		Password: password,
	})

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return &apiClient{cli}, nil 
}