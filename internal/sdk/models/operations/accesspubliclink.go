// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AccessPublicLinkRequest struct {
	Filename string `pathParam:"style=simple,explode=false,name=filename"`
	ID       string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AccessPublicLinkRequest) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *AccessPublicLinkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AccessPublicLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AccessPublicLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccessPublicLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccessPublicLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
