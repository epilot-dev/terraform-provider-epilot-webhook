# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "4yBbd5fpsHrHvMJ7kQeHJJ"
resource "epilot-webhook_webhook" "my_epilot-webhook_webhook" {
  auth = {
    api_key_config    = null
    auth_type         = "OAUTH_CLIENT_CREDENTIALS"
    basic_auth_config = null
    oauth_config = {
      client_id     = "92335773-69c3-4310-9cd6-a7b8a99f2e64"
      client_secret = null
      custom_parameter_list = [
        {
          key   = "scope"
          type  = "body"
          value = "target-entity:92335773-69c3-4310-9cd6-a7b8a99f2e64"
        },
      ]
      endpoint    = "https://iam.staging.wanyplace.de/oauth2/token"
      http_method = "POST"
    }
  }
  creation_time    = "2026-01-26T15:30:14.835Z"
  enable_static_ip = null
  enabled          = true
  event_name       = "erp_synchronization"
  filter = {
    key_to_filter    = "metadata.webhook_id"
    supported_values = ["4yBbd5fpsHrHvMJ7kQeHJJ"]
  }
  filter_conditions  = null
  http_method        = "PUT"
  jsonata_expression = "{\n  \"acknowledgeId\": entity.ack_id,\n  \"context\": {\n    \"identityId\": entity.identification.\"contact.identity_id\",\n    \"customerAccountId\": entity.identification.\"billing_account.external_id\"\n  },\n  \"data\": (\n    $payment := entity.latest.payment_method[\n      type = \"payment_sepa\" and (\n        (\"CREDIT\" in _tags and $not(\"DEBIT\" in _tags)) or \n        (\"DEBIT\" in _tags and \"CREDIT\" in _tags)\n      )\n    ][0];\n    $payment ? {\n      \"iban\": $payment.data.iban,\n      \"accountHolder\": $payment.data.fullname,\n      \"bankAccountUsage\": (\n        \"DEBIT\" in $payment._tags and \"CREDIT\" in $payment._tags \n          ? \"DEBIT_AND_CREDIT\" \n          : \"CREDIT\"\n      )\n    } : {}\n  )\n}"
  manifest           = ["d98d25e7-5d1e-4272-8b7b-7e92234e50d7"]
  name               = "Service - ERP - Änderung Bankverbindung / SEPA - Hinzufügen Bankverbindung (Entität: Rechnungseinheit)"
  payload_configuration = {
    custom_headers = {
      Ocp-Apim-Subscription-Key = "0b292c6f0ed8428ab86dc787e5cdfa41"
    }
    hydrate_entity             = false
    include_activity           = false
    include_changed_attributes = false
    include_relations          = false
  }
  status = "active"
  url    = "https://api.one2one.wilken.de/epilot/s3lf/v1/account/payment-method/bank"
}
