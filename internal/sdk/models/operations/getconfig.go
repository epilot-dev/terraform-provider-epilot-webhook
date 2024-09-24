// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"net/http"
)

type GetConfigRequest struct {
	// Short uuid to identify the webhook configuration.
	ConfigID string `pathParam:"style=simple,explode=false,name=configId"`
}

func (o *GetConfigRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

type GetConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// No configuration found for this id
	ErrorResp *shared.ErrorResp
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success - if the config is updated successfully
	WebhookConfig *shared.WebhookConfig
}

func (o *GetConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConfigResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *GetConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConfigResponse) GetWebhookConfig() *shared.WebhookConfig {
	if o == nil {
		return nil
	}
	return o.WebhookConfig
}
