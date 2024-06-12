// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type UpdateTaskRequestBody struct {
	// Task content. This value may contain markdown-formatted text and hyperlinks.
	Content *string `json:"content,omitempty"`
	// A description for the task. This value may contain markdown-formatted text and hyperlinks.
	Description *string `json:"description,omitempty"`
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
	// The responsible user ID or null to unset (for shared tasks).
	AssigneeID *string `json:"assignee_id,omitempty"`
	// A positive integer for the task duration, or null to unset. If specified, you must define a duration_unit.
	DurationAmount *int64 `json:"duration,omitempty"`
	// The unit of time for the duration. Must be either 'minute' or 'day'.
	DurationUnit *string `json:"duration_unit,omitempty"`
}

func (o *UpdateTaskRequestBody) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *UpdateTaskRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateTaskRequestBody) GetLabels() []string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *UpdateTaskRequestBody) GetPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *UpdateTaskRequestBody) GetDueString() *string {
	if o == nil {
		return nil
	}
	return o.DueString
}

func (o *UpdateTaskRequestBody) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *UpdateTaskRequestBody) GetDueDatetime() *string {
	if o == nil {
		return nil
	}
	return o.DueDatetime
}

func (o *UpdateTaskRequestBody) GetDueLang() *string {
	if o == nil {
		return nil
	}
	return o.DueLang
}

func (o *UpdateTaskRequestBody) GetAssigneeID() *string {
	if o == nil {
		return nil
	}
	return o.AssigneeID
}

func (o *UpdateTaskRequestBody) GetDurationAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.DurationAmount
}

func (o *UpdateTaskRequestBody) GetDurationUnit() *string {
	if o == nil {
		return nil
	}
	return o.DurationUnit
}

type UpdateTaskRequest struct {
	// The ID of the task to update.
	TaskID      string                `pathParam:"style=simple,explode=false,name=taskId"`
	RequestBody UpdateTaskRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateTaskRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *UpdateTaskRequest) GetRequestBody() UpdateTaskRequestBody {
	if o == nil {
		return UpdateTaskRequestBody{}
	}
	return o.RequestBody
}

type UpdateTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Task updated.
	Task *shared.Task
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *UpdateTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTaskResponse) GetTask() *shared.Task {
	if o == nil {
		return nil
	}
	return o.Task
}

func (o *UpdateTaskResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}
