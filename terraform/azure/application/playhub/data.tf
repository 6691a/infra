data "terraform_remote_state" "communication" {
  backend = "azurerm"
  config = {
    resource_group_name = "playhub"
    storage_account_name = "playhubs"
    container_name       = "terraform"
    key                  = "communication/playhub.tfstate"
  }
}