// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OAuthConfig struct {
	ClientID            types.String           `tfsdk:"client_id"`
	ClientSecret        types.String           `tfsdk:"client_secret"`
	CustomParameterList []CustomOAuthParameter `tfsdk:"custom_parameter_list"`
	Endpoint            types.String           `tfsdk:"endpoint"`
	HTTPMethod          types.String           `tfsdk:"http_method"`
}
