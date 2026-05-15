# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform
resource "epilot-webhook_webhook" "my_epilot-webhook_webhook" {
  creation_time    = "2026-04-24T14:26:16.933Z"
  delivery_mode    = null
  enable_static_ip = null
  enabled          = true
  event_name       = "automation_flow_target"
  filter = {
    key_to_filter    = "metadata.webhook_id"
    supported_values = ["iwpsEAgVqvnFuvqkiDkYZC"]
  }
  filter_conditions  = null
  http_method        = "POST"
  jsonata_expression = null
  manifest           = ["ff25ba9a-961f-4035-8304-dc57a9735e0c"]
  multipart_config   = null
  name               = "GetAnschlussobjekt_IS-U not created by speakeasy"
  payload_configuration = {
    apply_changesets           = null
    custom_headers             = null
    hydrate_entity             = false
    include_activity           = false
    include_changed_attributes = false
    include_relations          = false
  }
  protected    = true
  secure_proxy = null
  status       = "incomplete"
  url          = "https://epilot.apis.countandcare.de/api/v1/connection-object"
}
