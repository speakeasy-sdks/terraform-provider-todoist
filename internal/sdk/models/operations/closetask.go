// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type CloseTaskRequest struct {
	// The ID of the task to close.
	TaskID int64 `pathParam:"style=simple,explode=false,name=taskId"`
}

func (o *CloseTaskRequest) GetTaskID() int64 {
	if o == nil {
		return 0
	}
	return o.TaskID
}

type CloseTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *CloseTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CloseTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CloseTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CloseTaskResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}