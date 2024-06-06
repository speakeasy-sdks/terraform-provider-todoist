// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type UpdateSectionRequest struct {
	// Section ID.
	SectionID string `pathParam:"style=simple,explode=false,name=sectionId"`
	// Section name.
	Name string `queryParam:"style=form,explode=true,name=name"`
}

func (o *UpdateSectionRequest) GetSectionID() string {
	if o == nil {
		return ""
	}
	return o.SectionID
}

func (o *UpdateSectionRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type UpdateSectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Section name updated.
	Section *shared.Section
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *UpdateSectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSectionResponse) GetSection() *shared.Section {
	if o == nil {
		return nil
	}
	return o.Section
}

func (o *UpdateSectionResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}
