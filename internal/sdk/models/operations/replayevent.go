// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"net/http"
)

type ReplayEventRequest struct {
	// Short uuid to identify the webhook configuration.
	ConfigID string `pathParam:"style=simple,explode=false,name=configId"`
	// Event id
	EventID string `pathParam:"style=simple,explode=false,name=eventId"`
}

func (o *ReplayEventRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

func (o *ReplayEventRequest) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

type ReplayEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// No events found
	ErrorResp *shared.ErrorResp
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReplayEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplayEventResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *ReplayEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplayEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
