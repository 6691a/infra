terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.86.0"
    }
  }
  backend "azurerm" {
    resource_group_name  = "playhub"
    storage_account_name = "playhubs"
    container_name       = "terraform"
    key                  = "resource_group/playhub.tfstate"
  }
}

provider "azurerm" {
  features {}
}
