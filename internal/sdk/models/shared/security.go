// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	APIKey *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *Security) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}
