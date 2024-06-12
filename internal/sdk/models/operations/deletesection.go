// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type DeleteSectionRequest struct {
	// Section ID.
	SectionID string `pathParam:"style=simple,explode=false,name=sectionId"`
}

func (o *DeleteSectionRequest) GetSectionID() string {
	if o == nil {
		return ""
	}
	return o.SectionID
}

type DeleteSectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *DeleteSectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteSectionResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}