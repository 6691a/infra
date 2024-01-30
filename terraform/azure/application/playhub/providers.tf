terraform {
  required_providers {
    azuread = {
      source = "hashicorp/azuread"
      version = "2.47.0"
    }
     azurerm = {
      source = "hashicorp/azurerm"
      version = "3.86.0"
    }
  }
  backend "azurerm" {
    resource_group_name  = "playhub"
    storage_account_name = "playhubs"
    container_name       = "terraform"
    key                  = "application/playhub.tfstate"
  }
}

provider "azuread" {
}

provider "azurerm" {
  features {}
}