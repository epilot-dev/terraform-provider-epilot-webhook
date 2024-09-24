// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"net/http"
)

type UploadFileRequest struct {
	UploadFilePayload *shared.UploadFilePayload `request:"mediaType=application/json"`
	// file entity id
	FileEntityID *string `queryParam:"style=form,explode=true,name=file_entity_id"`
}

func (o *UploadFileRequest) GetUploadFilePayload() *shared.UploadFilePayload {
	if o == nil {
		return nil
	}
	return o.UploadFilePayload
}

func (o *UploadFileRequest) GetFileEntityID() *string {
	if o == nil {
		return nil
	}
	return o.FileEntityID
}

type S3ref struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func (o *S3ref) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *S3ref) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

// UploadFileResponseBody - Pre-signed URL for POST / PUT upload
type UploadFileResponseBody struct {
	// Returned only if file is permanent i.e. file_entity_id is passed
	PublicURL *string `json:"public_url,omitempty"`
	S3ref     *S3ref  `json:"s3ref,omitempty"`
	UploadURL *string `json:"upload_url,omitempty"`
}

func (o *UploadFileResponseBody) GetPublicURL() *string {
	if o == nil {
		return nil
	}
	return o.PublicURL
}

func (o *UploadFileResponseBody) GetS3ref() *S3ref {
	if o == nil {
		return nil
	}
	return o.S3ref
}

func (o *UploadFileResponseBody) GetUploadURL() *string {
	if o == nil {
		return nil
	}
	return o.UploadURL
}

type UploadFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Pre-signed URL for POST / PUT upload
	Object *UploadFileResponseBody
}

func (o *UploadFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadFileResponse) GetObject() *UploadFileResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
