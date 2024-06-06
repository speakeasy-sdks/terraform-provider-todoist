// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type CreateSectionRequest struct {
	// Project ID this section should belong to
	ProjectID string `queryParam:"style=form,explode=true,name=project_id"`
	// Section name
	Name string `queryParam:"style=form,explode=true,name=name"`
	// Order among other sections in a project
	Order *int64 `queryParam:"style=form,explode=true,name=order"`
}

func (o *CreateSectionRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *CreateSectionRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateSectionRequest) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

type CreateSectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Section created.
	Section *shared.Section
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *CreateSectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSectionResponse) GetSection() *shared.Section {
	if o == nil {
		return nil
	}
	return o.Section
}

func (o *CreateSectionResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}