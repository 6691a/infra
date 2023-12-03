variable "root_path" {
  type    = string
  default = "/Users/kimjun/Documents/projects/infra/packer/proxmox"
}

locals {
  credentials           = yamldecode(file(format("%s/credentials.yaml", var.root_path)))
  settings              = yamldecode(file("./settings.yaml"))
  vm                    = local.settings.vm
  disk                  = local.settings.disk
  ssh                   = local.settings.ssh
  node_exporter_version = "1.6.0"
}

