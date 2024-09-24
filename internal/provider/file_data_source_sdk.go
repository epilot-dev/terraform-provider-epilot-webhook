// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-webhook/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *FileDataSourceModel) RefreshFromSharedFileEntity(resp *shared.FileEntity) {
	if resp != nil {
		if len(resp.Additional) > 0 {
			r.Additional = make(map[string]types.String)
			for key, value := range resp.Additional {
				result, _ := json.Marshal(value)
				r.Additional[key] = types.StringValue(string(result))
			}
		}
		if resp.ACL == nil {
			r.ACL = nil
		} else {
			r.ACL = &tfTypes.BaseEntityACL{}
			r.ACL.Delete = []types.String{}
			for _, v := range resp.ACL.Delete {
				r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
			}
			r.ACL.Edit = []types.String{}
			for _, v := range resp.ACL.Edit {
				r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
			}
			r.ACL.View = []types.String{}
			for _, v := range resp.ACL.View {
				r.ACL.View = append(r.ACL.View, types.StringValue(v))
			}
		}
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		r.Org = types.StringPointerValue(resp.Org)
		r.Owners = []tfTypes.BaseEntityOwner{}
		if len(r.Owners) > len(resp.Owners) {
			r.Owners = r.Owners[:len(resp.Owners)]
		}
		for ownersCount, ownersItem := range resp.Owners {
			var owners1 tfTypes.BaseEntityOwner
			owners1.OrgID = types.StringValue(ownersItem.OrgID)
			owners1.UserID = types.StringPointerValue(ownersItem.UserID)
			if ownersCount+1 > len(r.Owners) {
				r.Owners = append(r.Owners, owners1)
			} else {
				r.Owners[ownersCount].OrgID = owners1.OrgID
				r.Owners[ownersCount].UserID = owners1.UserID
			}
		}
		r.Purpose = []types.String{}
		for _, v := range resp.Purpose {
			r.Purpose = append(r.Purpose, types.StringValue(v))
		}
		if resp.Schema != nil {
			r.Schema = types.StringValue(string(*resp.Schema))
		} else {
			r.Schema = types.StringNull()
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Title = types.StringPointerValue(resp.Title)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
		if resp.AccessControl != nil {
			r.AccessControl = types.StringValue(string(*resp.AccessControl))
		} else {
			r.AccessControl = types.StringNull()
		}
		r.CustomDownloadURL = types.StringPointerValue(resp.CustomDownloadURL)
		r.Filename = types.StringPointerValue(resp.Filename)
		r.ID = types.StringPointerValue(resp.ID)
		r.MimeType = types.StringPointerValue(resp.MimeType)
		r.PublicURL = types.StringPointerValue(resp.PublicURL)
		r.ReadableSize = types.StringPointerValue(resp.ReadableSize)
		if resp.S3ref == nil {
			r.S3ref = nil
		} else {
			r.S3ref = &tfTypes.S3Ref{}
			r.S3ref.Bucket = types.StringValue(resp.S3ref.Bucket)
			r.S3ref.Key = types.StringValue(resp.S3ref.Key)
		}
		r.SizeBytes = types.Int64PointerValue(resp.SizeBytes)
		r.SourceURL = types.StringPointerValue(resp.SourceURL)
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
		r.Versions = []tfTypes.FileItem{}
		if len(r.Versions) > len(resp.Versions) {
			r.Versions = r.Versions[:len(resp.Versions)]
		}
		for versionsCount, versionsItem := range resp.Versions {
			var versions1 tfTypes.FileItem
			versions1.Filename = types.StringPointerValue(versionsItem.Filename)
			versions1.MimeType = types.StringPointerValue(versionsItem.MimeType)
			versions1.ReadableSize = types.StringPointerValue(versionsItem.ReadableSize)
			if versionsItem.S3ref == nil {
				versions1.S3ref = nil
			} else {
				versions1.S3ref = &tfTypes.S3Ref{}
				versions1.S3ref.Bucket = types.StringValue(versionsItem.S3ref.Bucket)
				versions1.S3ref.Key = types.StringValue(versionsItem.S3ref.Key)
			}
			versions1.SizeBytes = types.Int64PointerValue(versionsItem.SizeBytes)
			if versionsCount+1 > len(r.Versions) {
				r.Versions = append(r.Versions, versions1)
			} else {
				r.Versions[versionsCount].Filename = versions1.Filename
				r.Versions[versionsCount].MimeType = versions1.MimeType
				r.Versions[versionsCount].ReadableSize = versions1.ReadableSize
				r.Versions[versionsCount].S3ref = versions1.S3ref
				r.Versions[versionsCount].SizeBytes = versions1.SizeBytes
			}
		}
	}
}