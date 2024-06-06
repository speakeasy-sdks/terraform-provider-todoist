// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ErrorModel struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func (o *ErrorModel) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ErrorModel) GetCode() int64 {
	if o == nil {
		return 0
	}
	return o.Code
}