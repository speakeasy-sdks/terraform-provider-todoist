// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-todoist/internal/provider/types"
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"time"
)

func (r *TaskDataSourceModel) RefreshFromSharedTask(resp *shared.Task) {
	if resp != nil {
		r.AssigneeID = types.StringPointerValue(resp.AssigneeID)
		r.AssignerID = types.StringPointerValue(resp.AssignerID)
		r.CommentCount = types.Int64PointerValue(resp.CommentCount)
		r.Content = types.StringPointerValue(resp.Content)
		r.CreatedAt = types.StringPointerValue(resp.CreatedAt)
		r.CreatorID = types.StringPointerValue(resp.CreatorID)
		r.Description = types.StringPointerValue(resp.Description)
		if resp.Due == nil {
			r.Due = nil
		} else {
			r.Due = &tfTypes.TaskDue{}
			if resp.Due.DueObject == nil {
				r.Due.DueObject = nil
			} else {
				r.Due.DueObject = &tfTypes.Due{}
				r.Due.DueObject.Date = types.StringValue(resp.Due.DueObject.Date.String())
				if resp.Due.DueObject.Datetime != nil {
					r.Due.DueObject.Datetime = types.StringValue(resp.Due.DueObject.Datetime.Format(time.RFC3339Nano))
				} else {
					r.Due.DueObject.Datetime = types.StringNull()
				}
				r.Due.DueObject.IsRecurring = types.BoolValue(resp.Due.DueObject.IsRecurring)
				r.Due.DueObject.String = types.StringValue(resp.Due.DueObject.String)
				r.Due.DueObject.Timezone = types.StringPointerValue(resp.Due.DueObject.Timezone)
			}
		}
		if resp.Duration == nil {
			r.Duration = nil
		} else {
			r.Duration = &tfTypes.TaskDuration{}
			if resp.Duration.DurationObject == nil {
				r.Duration.DurationObject = nil
			} else {
				r.Duration.DurationObject = &tfTypes.Duration{}
				r.Duration.DurationObject.Amount = types.Int64PointerValue(resp.Duration.DurationObject.Amount)
				r.Duration.DurationObject.Unit = types.StringPointerValue(resp.Duration.DurationObject.Unit)
			}
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.IsCompleted = types.BoolPointerValue(resp.IsCompleted)
		r.Labels = []types.String{}
		for _, v := range resp.Labels {
			r.Labels = append(r.Labels, types.StringValue(v))
		}
		r.Order = types.Int64PointerValue(resp.Order)
		r.ParentID = types.StringPointerValue(resp.ParentID)
		r.Priority = types.Int64PointerValue(resp.Priority)
		r.ProjectID = types.StringPointerValue(resp.ProjectID)
		r.SectionID = types.StringPointerValue(resp.SectionID)
		r.URL = types.StringPointerValue(resp.URL)
	}
}
