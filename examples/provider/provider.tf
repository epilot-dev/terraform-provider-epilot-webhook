terraform {
  required_providers {
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.4.0"
    }
  }
}

provider "epilot-webhook" {
  # Configuration options
}
