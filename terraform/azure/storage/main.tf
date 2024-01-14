resource "azurerm_storage_account" "storage" {
  name                     = var.storage_name
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_storage_container" "private_container" {
    name                  = var.private_container_name
    storage_account_name  = azurerm_storage_account.storage.name
    container_access_type = "private"
}

resource "azurerm_storage_container" "blob_container" {
    name                  = var.blob_container_name
    storage_account_name  = azurerm_storage_account.storage.name
    container_access_type = "blob"
}