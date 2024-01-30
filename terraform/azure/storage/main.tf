resource "azurerm_storage_account" "storage" {
  name                     = var.storage_name
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  blob_properties {
      cors_rule {
      allowed_headers = ["*"]
      allowed_methods = ["GET", "HEAD", "OPTIONS", "PUT"]
      allowed_origins = ["https://playhub.kr", "http://localhost:8000", "http://127.0.0.1:8000"]
      exposed_headers = ["*"]
      max_age_in_seconds = 3600
    }
  }
}

resource "azurerm_storage_container" "private_container" {
    name                  = "private"
    storage_account_name  = azurerm_storage_account.storage.name
    container_access_type = "private"
}

resource "azurerm_storage_container" "media_container" {
    name                  = "media"
    storage_account_name  = azurerm_storage_account.storage.name
    container_access_type = "private"
}

resource "azurerm_storage_container" "static_container" {
    name                  = "static"
    storage_account_name  = azurerm_storage_account.storage.name
    container_access_type = "blob"
}

resource "azurerm_storage_container" "terraform" {
    name                  = "terraform"
    storage_account_name  = azurerm_storage_account.storage.name
    container_access_type = "private"
}
