// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"net/http"
)

type GetFileRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// When passed true, the response will contain only fields that match the schema, with non-matching fields included in `__additional`
	Strict *bool `default:"false" queryParam:"style=form,explode=true,name=strict"`
}

func (g GetFileRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetFileRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetFileRequest) GetStrict() *bool {
	if o == nil {
		return nil
	}
	return o.Strict
}

type GetFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// File Entity
	FileEntity *shared.FileEntity
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFileResponse) GetFileEntity() *shared.FileEntity {
	if o == nil {
		return nil
	}
	return o.FileEntity
}

func (o *GetFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
