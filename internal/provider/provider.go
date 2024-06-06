// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

var _ provider.Provider = &TodoistProvider{}

type TodoistProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// TodoistProviderModel describes the provider data model.
type TodoistProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
	APIKey    types.String `tfsdk:"api_key"`
}

func (p *TodoistProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "todoist"
	resp.Version = p.version
}

func (p *TodoistProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `Todoist REST API v2 for ChatGPT Actions`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.todoist.com/rest/v2)",
				Optional:            true,
				Required:            false,
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *TodoistProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data TodoistProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.todoist.com/rest/v2"
	}

	apiKey := new(string)
	if !data.APIKey.IsUnknown() && !data.APIKey.IsNull() {
		*apiKey = data.APIKey.ValueString()
	} else {
		apiKey = nil
	}
	security := shared.Security{
		APIKey: apiKey,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(http.DefaultClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *TodoistProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectResource,
		NewTaskResource,
	}
}

func (p *TodoistProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProjectDataSource,
		NewTaskDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TodoistProvider{
			version: version,
		}
	}
}
