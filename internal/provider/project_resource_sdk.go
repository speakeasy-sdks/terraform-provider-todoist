// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
)

func (r *ProjectResourceModel) RefreshFromSharedProject(resp *shared.Project) {
	if resp != nil {
		r.Color = types.StringPointerValue(resp.Color)
		r.CommentCount = types.Int64PointerValue(resp.CommentCount)
		r.ID = types.StringPointerValue(resp.ID)
		r.IsInboxProject = types.BoolPointerValue(resp.IsInboxProject)
		r.IsShared = types.BoolPointerValue(resp.IsShared)
		r.IsTeamInbox = types.BoolPointerValue(resp.IsTeamInbox)
		r.IsFavorite = types.BoolPointerValue(resp.IsFavorite)
		r.Name = types.StringPointerValue(resp.Name)
		r.Order = types.Int64PointerValue(resp.Order)
		r.ParentID = types.StringPointerValue(resp.ParentID)
		r.URL = types.StringPointerValue(resp.URL)
		r.ViewStyle = types.StringPointerValue(resp.ViewStyle)
	}
}
