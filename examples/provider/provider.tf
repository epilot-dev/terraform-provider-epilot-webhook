terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.5.2"
    }
  }
}

provider "epilot-webhook" {
  # Configuration options
}