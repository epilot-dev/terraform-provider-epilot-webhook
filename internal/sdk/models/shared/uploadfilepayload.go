// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/internal/utils"
)

type UploadFilePayload struct {
	Filename string `json:"filename"`
	// Used to index the file at the storage layer, which helps when browsing for this file
	IndexTag *string `json:"index_tag,omitempty"`
	// Allows passing in custom metadata for the file, expects key-value pairs of string type
	Metadata map[string]string `json:"metadata,omitempty"`
	// MIME type of file
	MimeType *string `default:"application/octet-stream" json:"mime_type"`
}

func (u UploadFilePayload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UploadFilePayload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UploadFilePayload) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *UploadFilePayload) GetIndexTag() *string {
	if o == nil {
		return nil
	}
	return o.IndexTag
}

func (o *UploadFilePayload) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *UploadFilePayload) GetMimeType() *string {
	if o == nil {
		return nil
	}
	return o.MimeType
}
