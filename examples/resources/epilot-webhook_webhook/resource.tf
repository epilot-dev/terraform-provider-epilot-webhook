# resource "epilot-webhook_webhook" "my_webhook" {
#   auth = {
#     api_key_config = {
#       key_name  = "...my_key_name..."
#       key_value = "...my_key_value..."
#     }
#     auth_type = "NONE"
#     basic_auth_config = {
#       password = "...my_password..."
#       username = "...my_username..."
#     }
#     oauth_config = {
#       client_id     = "...my_client_id..."
#       client_secret = "...my_client_secret..."
#       custom_parameter_list = [
#         {
#           key   = "...my_key..."
#           type  = "query"
#           value = "...my_value..."
#         }
#       ]
#       endpoint    = "...my_endpoint..."
#       http_method = "PUT"
#     }
#   }
#   creation_time    = "2021-04-27T12:01:13.000Z"
#   enable_static_ip = false
#   enabled          = false
#   event_name       = "...my_event_name..."
#   filter = {
#     key_to_filter = "...my_key_to_filter..."
#     supported_values = [
#       "..."
#     ]
#   }
#   filter_conditions = {
#     conditions = [
#       {
#         field          = "...my_field..."
#         field_type     = "number"
#         is_array_field = true
#         operation      = "none_of"
#         values = [
#           "..."
#         ]
#       }
#     ]
#     logical_operator = "AND"
#   }
#   http_method        = "GET"
#   jsonata_expression = "...my_jsonata_expression..."
#   manifest = [
#     "123e4567-e89b-12d3-a456-426614174000"
#   ]
#   name = "...my_name..."
#   payload_configuration = {
#     custom_headers = {
#       key = "value"
#     }
#     hydrate_entity             = true
#     include_activity           = false
#     include_changed_attributes = true
#     include_relations          = true
#   }
#   status = "active"
#   url    = "...my_url..."
# }