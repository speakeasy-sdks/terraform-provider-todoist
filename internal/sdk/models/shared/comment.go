// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-todoist/internal/sdk/internal/utils"
	"time"
)

// Attachment file metadata (will be null if there is no attachment). Format varies depending on the type of attachment, as detailed in the Sync API documentation.
type Attachment struct {
}

type Comment struct {
	// Comment ID.
	ID *string `json:"id,omitempty"`
	// Comment's task ID (will be null if the comment belongs to a project).
	TaskID *string `json:"task_id,omitempty"`
	// Comment's project ID (will be null if the comment belongs to a task).
	ProjectID *string `json:"project_id,omitempty"`
	// Date and time when comment was added, in RFC3339 format in UTC.
	PostedAt *time.Time `json:"posted_at,omitempty"`
	// Comment content. This value may contain markdown-formatted text and hyperlinks. Details on markdown support can be found in the Text Formatting article in the Help Center.
	Content *string `json:"content,omitempty"`
	// Attachment file metadata (will be null if there is no attachment). Format varies depending on the type of attachment, as detailed in the Sync API documentation.
	Attachment *Attachment `json:"attachment,omitempty"`
}

func (c Comment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Comment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Comment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Comment) GetTaskID() *string {
	if o == nil {
		return nil
	}
	return o.TaskID
}

func (o *Comment) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *Comment) GetPostedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PostedAt
}

func (o *Comment) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *Comment) GetAttachment() *Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}