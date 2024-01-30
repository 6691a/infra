data "terraform_remote_state" "resource_group" {
  backend = "azurerm"
  config = {
    resource_group_name = "playhub"
    storage_account_name = "playhubs"
    container_name       = "terraform"
    key                  = "resource_group/playhub.tfstate"
  }
}

