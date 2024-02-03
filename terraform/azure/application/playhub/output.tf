output "client_secret" {
  description = "Client Secret"
  value       = nonsensitive(azuread_application_password.password.value)
}