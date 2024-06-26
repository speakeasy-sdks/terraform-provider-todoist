// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type GetPersonalLabelRequest struct {
	// The ID of the label to retrieve.
	LabelID int64 `pathParam:"style=simple,explode=false,name=label_id"`
}

func (o *GetPersonalLabelRequest) GetLabelID() int64 {
	if o == nil {
		return 0
	}
	return o.LabelID
}

type GetPersonalLabelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Label details.
	Label *shared.Label
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *GetPersonalLabelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPersonalLabelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPersonalLabelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPersonalLabelResponse) GetLabel() *shared.Label {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *GetPersonalLabelResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}
