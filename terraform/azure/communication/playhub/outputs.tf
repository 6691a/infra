output "smtp_role" {
    value = azurerm_role_definition.smtp_role
}

output "communication" {
    value = azurerm_communication_service.playhub
}