terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.6.1"
    }
  }
}

provider "epilot-webhook" {
  # Configuration options
}