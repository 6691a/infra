terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.82.0"
    }
  }
}

provider "google" {
  credentials = file(local.credentials_path)
  project     = local.project
  region      = var.region
}


resource "google_artifact_registry_repository" "repository" {
  repository_id  = var.repository_name
  location       = var.region
  format         = "DOCKER"
}