// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"net/http"
)

type GetFailuresForConfigRequest struct {
	// Short uuid to identify the webhook configuration.
	ConfigID string `pathParam:"style=simple,explode=false,name=configId"`
	// Optional Key. To be provided when loading paginated data. It is always returned initially by this API, if pagination is needed.
	LastLoadedEventID   *string `queryParam:"style=form,explode=true,name=lastLoadedEventId"`
	LastLoadedTimestamp *string `queryParam:"style=form,explode=true,name=lastLoadedTimestamp"`
}

func (o *GetFailuresForConfigRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

func (o *GetFailuresForConfigRequest) GetLastLoadedEventID() *string {
	if o == nil {
		return nil
	}
	return o.LastLoadedEventID
}

func (o *GetFailuresForConfigRequest) GetLastLoadedTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.LastLoadedTimestamp
}

type GetFailuresForConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// No configuration found for this id
	ErrorResp *shared.ErrorResp
	// Success
	FailuresResp *shared.FailuresResp
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetFailuresForConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFailuresForConfigResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *GetFailuresForConfigResponse) GetFailuresResp() *shared.FailuresResp {
	if o == nil {
		return nil
	}
	return o.FailuresResp
}

func (o *GetFailuresForConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFailuresForConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
