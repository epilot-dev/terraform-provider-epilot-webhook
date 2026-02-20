terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.9.0"
    }
  }
}

provider "epilot-webhook" {
  server_url = "..." # Optional
}