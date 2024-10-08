---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-webhook_webhook Resource - terraform-provider-epilot-webhook"
subcategory: ""
description: |-
  Webhook Resource
---

# epilot-webhook_webhook (Resource)

Webhook Resource

## Example Usage

```terraform
resource "epilot-webhook_webhook" "my_webhook" {
  event_name       = "automation_flow_target"
  name        = "Generated by Terraform"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `event_name` (String)
- `name` (String)

### Optional

- `auth` (Attributes) (see [below for nested schema](#nestedatt--auth))
- `creation_time` (String) creation timestamp
- `enable_static_ip` (Boolean)
- `enabled` (Boolean)
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `http_method` (String) must be one of ["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"]
- `payload_configuration` (Attributes) Configuration for the webhook payload (see [below for nested schema](#nestedatt--payload_configuration))
- `status` (String) must be one of ["active", "inactive", "incomplete"]
- `url` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--auth"></a>
### Nested Schema for `auth`

Optional:

- `api_key_config` (Attributes) To be sent only if authType is API_KEY (see [below for nested schema](#nestedatt--auth--api_key_config))
- `auth_type` (String) Not Null; must be one of ["BASIC", "OAUTH_CLIENT_CREDENTIALS", "API_KEY"]
- `basic_auth_config` (Attributes) To be sent only if authType is BASIC (see [below for nested schema](#nestedatt--auth--basic_auth_config))
- `oauth_config` (Attributes) To be sent only if authType is OAUTH_CLIENT_CREDENTIALS (see [below for nested schema](#nestedatt--auth--oauth_config))

<a id="nestedatt--auth--api_key_config"></a>
### Nested Schema for `auth.api_key_config`

Optional:

- `key_name` (String) Not Null
- `key_value` (String)


<a id="nestedatt--auth--basic_auth_config"></a>
### Nested Schema for `auth.basic_auth_config`

Optional:

- `password` (String)
- `username` (String) Not Null


<a id="nestedatt--auth--oauth_config"></a>
### Nested Schema for `auth.oauth_config`

Optional:

- `client_id` (String) Not Null
- `client_secret` (String)
- `custom_parameter_list` (Attributes List) (see [below for nested schema](#nestedatt--auth--oauth_config--custom_parameter_list))
- `endpoint` (String) Https Endpoint for authentication. Not Null
- `http_method` (String) Not Null; must be one of ["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"]

<a id="nestedatt--auth--oauth_config--custom_parameter_list"></a>
### Nested Schema for `auth.oauth_config.custom_parameter_list`

Optional:

- `key` (String) Not Null
- `type` (String) Not Null; must be one of ["body", "query", "header"]
- `value` (String) Not Null




<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Optional:

- `key_to_filter` (String) Not Null
- `supported_values` (List of String) Not Null


<a id="nestedatt--payload_configuration"></a>
### Nested Schema for `payload_configuration`

Optional:

- `custom_headers` (Map of String) Object representing custom headers as key-value pairs.
- `hydrate_entity` (Boolean)
- `include_activity` (Boolean)
- `include_changed_attributes` (Boolean)
- `include_relations` (Boolean)

## Import

Import is supported using the following syntax:

```shell
terraform import epilot-webhook_webhook.my_epilot-webhook_webhook ""
```
