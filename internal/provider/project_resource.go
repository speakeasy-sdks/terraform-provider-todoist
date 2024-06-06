// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/speakeasy/terraform-provider-todoist/internal/planmodifiers/stringplanmodifier"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProjectResource{}
var _ resource.ResourceWithImportState = &ProjectResource{}

func NewProjectResource() resource.Resource {
	return &ProjectResource{}
}

// ProjectResource defines the resource implementation.
type ProjectResource struct {
	client *sdk.SDK
}

// ProjectResourceModel describes the resource data model.
type ProjectResourceModel struct {
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

func (r *ProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Project Resource",
		Attributes: map[string]schema.Attribute{
			"color": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The color of the project icon. Refer to the name column in the Colors guide for more info. https://developer.todoist.com/guides/#colors`,
			},
			"comment_count": schema.Int64Attribute{
				Computed:    true,
				Description: `Number of project comments.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Project ID.`,
			},
			"is_favorite": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
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
				Required:    true,
				Description: `Name of the project`,
			},
			"order": schema.Int64Attribute{
				Computed:    true,
				Description: `Project position under the same parent (read-only, will be 0 for inbox and team inbox projects).`,
			},
			"parent_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Parent project ID. Requires replacement if changed. `,
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: `URL to access this project in the Todoist web or mobile applications.`,
			},
			"view_style": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A string value (either list or board, default is list). This determines the way the project is displayed within the Todoist clients.`,
			},
		},
	}
}

func (r *ProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ProjectResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()
	parentID := new(string)
	if !data.ParentID.IsUnknown() && !data.ParentID.IsNull() {
		*parentID = data.ParentID.ValueString()
	} else {
		parentID = nil
	}
	color := new(string)
	if !data.Color.IsUnknown() && !data.Color.IsNull() {
		*color = data.Color.ValueString()
	} else {
		color = nil
	}
	isFavorite := new(bool)
	if !data.IsFavorite.IsUnknown() && !data.IsFavorite.IsNull() {
		*isFavorite = data.IsFavorite.ValueBool()
	} else {
		isFavorite = nil
	}
	viewStyle := new(string)
	if !data.ViewStyle.IsUnknown() && !data.ViewStyle.IsNull() {
		*viewStyle = data.ViewStyle.ValueString()
	} else {
		viewStyle = nil
	}
	request := operations.CreateProjectRequest{
		Name:       name,
		ParentID:   parentID,
		Color:      color,
		IsFavorite: isFavorite,
		ViewStyle:  viewStyle,
	}
	res, err := r.client.Projects.CreateProject(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Project == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProject(res.Project)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ProjectResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

func (r *ProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ProjectResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := data.ID.ValueString()
	name := new(string)
	if !data.Name.IsUnknown() && !data.Name.IsNull() {
		*name = data.Name.ValueString()
	} else {
		name = nil
	}
	color := new(string)
	if !data.Color.IsUnknown() && !data.Color.IsNull() {
		*color = data.Color.ValueString()
	} else {
		color = nil
	}
	isFavorite := new(bool)
	if !data.IsFavorite.IsUnknown() && !data.IsFavorite.IsNull() {
		*isFavorite = data.IsFavorite.ValueBool()
	} else {
		isFavorite = nil
	}
	viewStyle := new(string)
	if !data.ViewStyle.IsUnknown() && !data.ViewStyle.IsNull() {
		*viewStyle = data.ViewStyle.ValueString()
	} else {
		viewStyle = nil
	}
	request := operations.UpdateProjectRequest{
		ProjectID:  projectID,
		Name:       name,
		Color:      color,
		IsFavorite: isFavorite,
		ViewStyle:  viewStyle,
	}
	res, err := r.client.Projects.UpdateProject(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Project == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProject(res.Project)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ProjectResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteProjectRequest{
		ProjectID: projectID,
	}
	res, err := r.client.Projects.DeleteProject(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *ProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}