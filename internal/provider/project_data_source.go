// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ProjectDataSource{}
var _ datasource.DataSourceWithConfigure = &ProjectDataSource{}

func NewProjectDataSource() datasource.DataSource {
	return &ProjectDataSource{}
}

// ProjectDataSource is the data source implementation.
type ProjectDataSource struct {
	client *sdk.SDK
}

// ProjectDataSourceModel describes the data model.
type ProjectDataSourceModel struct {
	Color          types.String `tfsdk:"color"`
	CommentCount   types.Int64  `tfsdk:"comment_count"`
	ID             types.String `tfsdk:"id"`
	IsFavorite     types.Bool   `tfsdk:"is_favorite"`
	IsInboxProject types.Bool   `tfsdk:"is_inbox_project"`
	IsShared       types.Bool   `tfsdk:"is_shared"`
	IsTeamInbox    types.Bool   `tfsdk:"is_team_inbox"`
	Name           types.String `tfsdk:"name"`
	Order          types.Int64  `tfsdk:"order"`
	ParentID       types.String `tfsdk:"parent_id"`
	URL            types.String `tfsdk:"url"`
	ViewStyle      types.String `tfsdk:"view_style"`
}

// Metadata returns the data source type name.
func (r *ProjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the data source.
func (r *ProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Project DataSource",

		Attributes: map[string]schema.Attribute{
			"color": schema.StringAttribute{
				Computed:    true,
				Description: `The color of the project icon. Refer to the name column in the Colors guide for more info.`,
			},
			"comment_count": schema.Int64Attribute{
				Computed:    true,
				Description: `Number of project comments.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `Project ID.`,
			},
			"is_favorite": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the project is a favorite (a true or false value).`,
			},
			"is_inbox_project": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the project is the user's Inbox (read-only).`,
			},
			"is_shared": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the project is shared (read-only, a true or false value).`,
			},
			"is_team_inbox": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the project is the Team Inbox (read-only).`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Project name.`,
			},
			"order": schema.Int64Attribute{
				Computed:    true,
				Description: `Project position under the same parent (read-only, will be 0 for inbox and team inbox projects).`,
			},
			"parent_id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of parent project (will be null for top-level projects).`,
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: `URL to access this project in the Todoist web or mobile applications.`,
			},
			"view_style": schema.StringAttribute{
				Computed:    true,
				Description: `A string value (either list or board). This determines the way the project is displayed within the Todoist clients.`,
			},
		},
	}
}

func (r *ProjectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ProjectDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	projectID := data.ID.ValueString()
	request := operations.GetProjectRequest{
		ProjectID: projectID,
	}
	res, err := r.client.Projects.GetProject(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Project == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProject(res.Project)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}