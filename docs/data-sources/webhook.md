---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-webhook_webhook Data Source - terraform-provider-epilot-webhook"
subcategory: ""
description: |-
  Webhook DataSource
---

# epilot-webhook_webhook (Data Source)

Webhook DataSource

## Example Usage

```terraform
data "epilot-webhook_webhook" "my_webhook" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `auth` (Attributes) (see [below for nested schema](#nestedatt--auth))
- `creation_time` (String) creation timestamp
- `enable_static_ip` (Boolean)
- `enabled` (Boolean)
- `event_name` (String)
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `http_method` (String)
- `id` (String) The ID of this resource.
- `jsonata_expression` (String) JSONata expression to transform the payload
- `manifest` (List of String) Manifest ID used to create/update the entity
- `name` (String)
- `payload_configuration` (Attributes) Configuration for the webhook payload (see [below for nested schema](#nestedatt--payload_configuration))
- `status` (String)
- `url` (String)

<a id="nestedatt--auth"></a>
### Nested Schema for `auth`

Read-Only:

- `api_key_config` (Attributes) To be sent only if authType is API_KEY (see [below for nested schema](#nestedatt--auth--api_key_config))
- `auth_type` (String)
- `basic_auth_config` (Attributes) To be sent only if authType is BASIC (see [below for nested schema](#nestedatt--auth--basic_auth_config))
- `oauth_config` (Attributes) To be sent only if authType is OAUTH_CLIENT_CREDENTIALS (see [below for nested schema](#nestedatt--auth--oauth_config))

<a id="nestedatt--auth--api_key_config"></a>
### Nested Schema for `auth.api_key_config`

Read-Only:

- `key_name` (String)
- `key_value` (String)


<a id="nestedatt--auth--basic_auth_config"></a>
### Nested Schema for `auth.basic_auth_config`

Read-Only:

- `password` (String)
- `username` (String)


<a id="nestedatt--auth--oauth_config"></a>
### Nested Schema for `auth.oauth_config`

Read-Only:

- `client_id` (String)
- `client_secret` (String)
- `custom_parameter_list` (Attributes List) (see [below for nested schema](#nestedatt--auth--oauth_config--custom_parameter_list))
- `endpoint` (String) Https Endpoint for authentication
- `http_method` (String)

<a id="nestedatt--auth--oauth_config--custom_parameter_list"></a>
### Nested Schema for `auth.oauth_config.custom_parameter_list`

Read-Only:

- `key` (String)
- `type` (String)
- `value` (String)




<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Read-Only:

- `key_to_filter` (String)
- `supported_values` (List of String)


<a id="nestedatt--payload_configuration"></a>
### Nested Schema for `payload_configuration`

Read-Only:

- `custom_headers` (Map of String) Object representing custom headers as key-value pairs.
- `hydrate_entity` (Boolean)
- `include_activity` (Boolean)
- `include_changed_attributes` (Boolean)
- `include_relations` (Boolean)
