terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "2.1.1"
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
  server_url  = "https://file.dev.sls.epilot.io"
}
