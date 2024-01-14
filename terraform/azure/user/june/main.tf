resource "azuread_user" "new_user" {
    user_principal_name = "${var.user_name}@playhub.kr"
    display_name = var.user_name
    password = var.password

    force_password_change = true
}

resource "azurerm_role_assignment" "new_user_role" {
    scope = "/subscriptions/${var.subscription_id}"
    role_definition_name = "Contributor"
    principal_id = azuread_user.new_user.object_id
}




