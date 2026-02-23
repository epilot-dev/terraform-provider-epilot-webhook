terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.9.1"
    }
  }
}

provider "epilot-webhook" {
  server_url = "..." # Optional
}