// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type CreateTaskRequestBody struct {
	// Task content. This value may contain markdown-formatted text and hyperlinks.
	Content *string `json:"content,omitempty"`
	// A description for the task. This value may contain markdown-formatted text and hyperlinks.
	Description *string `json:"description,omitempty"`
	// Task project ID. If not set, the task is put in the user's Inbox.
	ProjectID *string `json:"project_id,omitempty"`
	// ID of the section to put the task into.
	SectionID *string `json:"section_id,omitempty"`
	// Parent task ID.
	ParentID *string `json:"parent_id,omitempty"`
	// Non-zero integer value used by clients to sort tasks under the same parent.
	Order *int64 `json:"order,omitempty"`
	// The task's labels (a list of names that may represent either personal or shared labels).
	Labels []string `json:"labels,omitempty"`
	// Task priority from 1 (normal) to 4 (urgent).
	Priority *int64 `json:"priority,omitempty"`
	// Human-defined task due date (ex. "next Monday," "Tomorrow"). Value is set using local (not UTC) time.
	DueString *string `json:"due_string,omitempty"`
	// Specific date in YYYY-MM-DD format relative to the user's timezone.
	DueDate *string `json:"due_date,omitempty"`
	// Specific date and time in RFC3339 format in UTC.
	DueDatetime *string `json:"due_datetime,omitempty"`
	// 2-letter code specifying the language in case due_string is not written in English.
	DueLang *string `json:"due_lang,omitempty"`
	// The responsible user ID (only applies to shared tasks).
	AssigneeID *string `json:"assignee_id,omitempty"`
	// A positive (greater than zero) integer for the amount of duration_unit the task will take. If specified, you must define a duration_unit.
	Duration *int64 `json:"duration,omitempty"`
	// The unit of time that the duration field above represents. Must be either minute or day. If specified, duration must be defined as well.
	DurationUnit *string `json:"duration_unit,omitempty"`
}

func (o *CreateTaskRequestBody) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *CreateTaskRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateTaskRequestBody) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *CreateTaskRequestBody) GetSectionID() *string {
	if o == nil {
		return nil
	}
	return o.SectionID
}

func (o *CreateTaskRequestBody) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *CreateTaskRequestBody) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *CreateTaskRequestBody) GetLabels() []string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *CreateTaskRequestBody) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *CreateTaskRequestBody) GetDueString() *string {
	if o == nil {
		return nil
	}
	return o.DueString
}

func (o *CreateTaskRequestBody) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *CreateTaskRequestBody) GetDueDatetime() *string {
	if o == nil {
		return nil
	}
	return o.DueDatetime
}

func (o *CreateTaskRequestBody) GetDueLang() *string {
	if o == nil {
		return nil
	}
	return o.DueLang
}

func (o *CreateTaskRequestBody) GetAssigneeID() *string {
	if o == nil {
		return nil
	}
	return o.AssigneeID
}

func (o *CreateTaskRequestBody) GetDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *CreateTaskRequestBody) GetDurationUnit() *string {
	if o == nil {
		return nil
	}
	return o.DurationUnit
}

type CreateTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Task created.
	Task *shared.Task
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *CreateTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTaskResponse) GetTask() *shared.Task {
	if o == nil {
		return nil
	}
	return o.Task
}

func (o *CreateTaskResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}