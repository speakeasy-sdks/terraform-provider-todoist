// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/models/shared"
	"net/http"
)

type CreateProjectRequest struct {
	// Name of the project
	Name string `queryParam:"style=form,explode=true,name=name"`
	// Parent project ID
	ParentID *string `queryParam:"style=form,explode=true,name=parent_id"`
	// The color of the project icon. Refer to the name column in the Colors guide for more info. https://developer.todoist.com/guides/#colors
	Color *string `queryParam:"style=form,explode=true,name=color"`
	// Whether the project is a favorite (a true or false value).
	IsFavorite *bool `queryParam:"style=form,explode=true,name=is_favorite"`
	// A string value (either list or board, default is list). This determines the way the project is displayed within the Todoist clients.
	ViewStyle *string `queryParam:"style=form,explode=true,name=view_style"`
}

func (o *CreateProjectRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateProjectRequest) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *CreateProjectRequest) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *CreateProjectRequest) GetIsFavorite() *bool {
	if o == nil {
		return nil
	}
	return o.IsFavorite
}

func (o *CreateProjectRequest) GetViewStyle() *string {
	if o == nil {
		return nil
	}
	return o.ViewStyle
}

type CreateProjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Project created.
	Project *shared.Project
	// Unexpected error
	ErrorModel *shared.ErrorModel
}

func (o *CreateProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateProjectResponse) GetProject() *shared.Project {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *CreateProjectResponse) GetErrorModel() *shared.ErrorModel {
	if o == nil {
		return nil
	}
	return o.ErrorModel
}
