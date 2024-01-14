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
}

provider "azuread" {
}

provider "azurerm" {
  features {}
}