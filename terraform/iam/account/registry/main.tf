resource "google_service_account" "registry_user" {
  account_id   = "registry"
  display_name = "read only registry"
}

resource "google_service_account_key" "registry_user_key" {
  service_account_id = google_service_account.registry_user.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}

resource "local_file" "service_account_key_file" {
  content  = base64decode(google_service_account_key.registry_user_key.private_key)
  filename = "${path.module}/key.json"
}

resource "google_project_iam_custom_role" "read_only_registry" {
  role_id     = "read_only_registry"
  title       = "Read only registry"
  description = "Read only access to registry"
  permissions = [
    "artifactregistry.repositories.downloadArtifacts",
    "artifactregistry.repositories.list",
    "artifactregistry.repositories.get",
  ]
}

resource "google_project_iam_member" "registry_user_role" {
  project = local.credentials.project_id
  role    = "projects/${local.credentials.project_id}/roles/read_only_registry"
  member  = "serviceAccount:${google_service_account.registry_user.email}"
}
