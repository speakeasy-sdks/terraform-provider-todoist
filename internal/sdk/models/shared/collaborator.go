// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Collaborator struct {
	// Collaborator's user ID
	ID *string `json:"id,omitempty"`
	// Collaborator's name
	Name *string `json:"name,omitempty"`
	// Collaborator's email
	Email *string `json:"email,omitempty"`
}

func (o *Collaborator) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Collaborator) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Collaborator) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}