resource "azurerm_email_communication_service" "playhub" {
  name       = "playhub"
  resource_group_name = data.terraform_remote_state.resource_group.outputs.playhub.name
  data_location       = "Korea"
}

resource "azurerm_communication_service" "playhub" {
  name                = "playhub"
  resource_group_name = data.terraform_remote_state.resource_group.outputs.playhub.name
  data_location       = "Korea"
}

resource "azurerm_role_definition" "smtp_role" {
  name        = "SMTP"
  description = "SMTP role for Playhub"
  scope       = azurerm_communication_service.playhub.id

  permissions {
    actions     = [     
      "Microsoft.Communication/CommunicationServices/read", 
      "Microsoft.Communication/EmailServices/write",
    ]
    not_actions = []
  }
}
