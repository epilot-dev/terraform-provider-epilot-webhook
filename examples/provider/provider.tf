terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.5.1"
    }
  }
}

provider "epilot-webhook" {
  # Configuration options
}