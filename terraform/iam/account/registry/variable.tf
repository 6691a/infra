variable "root" {
    description = "The root path of the project"
    type = string
    default = "/Users/kimjun/Documents/projects/infra/secret/"
}

variable "credentials_path" {
  type = string
  default = "gcp_credential.json"
}

variable "region" {
  type = string
  default = "asia-northeast3"
}

locals {
  credentials_path = "${var.root}/${var.credentials_path}"
  credentials = jsondecode(file(local.credentials_path))
  project = local.credentials.project_id
}
