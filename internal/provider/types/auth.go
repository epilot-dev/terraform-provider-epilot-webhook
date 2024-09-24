// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Auth struct {
	APIKeyConfig    *APIKeyConfig    `tfsdk:"api_key_config"`
	AuthType        types.String     `tfsdk:"auth_type"`
	BasicAuthConfig *BasicAuthConfig `tfsdk:"basic_auth_config"`
	OauthConfig     *OAuthConfig     `tfsdk:"oauth_config"`
}