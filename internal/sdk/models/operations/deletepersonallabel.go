// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type DeletePersonalLabelRequest struct {
	// The ID of the label to delete.
	LabelID int64 `pathParam:"style=simple,explode=false,name=label_id"`
}

func (o *DeletePersonalLabelRequest) GetLabelID() int64 {
	if o == nil {
		return 0
	}
	return o.LabelID
}

type DeletePersonalLabelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *DeletePersonalLabelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePersonalLabelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePersonalLabelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeletePersonalLabelResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}
