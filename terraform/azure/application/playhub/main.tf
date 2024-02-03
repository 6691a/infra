resource "azuread_application" "playhub" {
    display_name   = "playhub"
    identifier_uris = ["https://playhub.kr"]
}

resource "azuread_service_principal" "playhub" {
    client_id = azuread_application.playhub.client_id
}

resource "azurerm_role_assignment" "communication_role" {
    scope = local.communication.id
    role_definition_id = local.smtp_role.role_definition_resource_id
    principal_id = azuread_service_principal.playhub.id
}

resource "azuread_application_password" "password" {
    application_id = azuread_application.playhub.id
    display_name = "playhub"
}