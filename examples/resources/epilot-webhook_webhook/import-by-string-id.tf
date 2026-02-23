import {
  to = epilot-webhook_webhook.my_epilot-webhook_webhook
  id = "4yBbd5fpsHrHvMJ7kQeHJJ"
}


terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.5.0"
    }
  }
}

variable "epilot_auth" {
  type        = string
  description = "epilot_auth"
}

provider "epilot-webhook" {
  # Configuration options
  epilot_auth = var.epilot_auth
}