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
				Type:        schema.TypeList,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENDPOINTS", []string{"localhost:2379"}),
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("username", "rootuser"),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("password", "rootuser"),
			},
			"is_auth_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("isAuthEnabled", true),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cluster_member_list_data_source": ClusterDataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"kv_resource": KvResource(),
		},
	}

	p.ConfigureContextFunc = configure

	return p

}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	*etcd.Client
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Setup a User-Agent for your API client (replace the provider name for yours):
	// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
	// TODO: myClient.UserAgent = userAgent
	var (
		diags diag.Diagnostics
		err   error
		cli   *etcd.Client
	)

	endpoints := d.Get("endpoints").([]string)

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	isAuthEnabled := d.Get("is_auth_enabled").(bool)

	if !isAuthEnabled {
		cli, err = etcd.New(etcd.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
		})

	} else {
		cli, err = etcd.New(etcd.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
			Username:    username,
			Password:    password,
		})
		cli.AuthEnable(ctx)
	}

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "",
			Detail:   "",
		})
		return nil, diags
	}

	defer cli.Close()

	return &apiClient{cli}, nil
}
