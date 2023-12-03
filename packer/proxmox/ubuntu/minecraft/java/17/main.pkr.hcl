packer {
  required_plugins {
    proxmox = {
      version = " >= 1.1.5"
      source  = "github.com/hashicorp/proxmox"
    }
  }
}

source "proxmox-clone" "ubuntu" {
  proxmox_url              = local.credentials.api_url
  username                 = local.credentials.api_token_id
  token                    = local.credentials.api_token_secret
  insecure_skip_tls_verify = true

  node        = local.vm.node
  vm_id       = local.vm.id
  vm_name     = local.vm.name
  clone_vm_id = 10000

  cores        = local.vm.cores
  memory       = local.vm.memory
  task_timeout = "10m"

  scsi_controller = "virtio-scsi-pci"

  network_adapters {
    bridge   = "vmbr0"
    model    = "virtio"
    firewall = "false"
  }

  ipconfig {
    ip      = local.network.ip
    gateway = local.network.gateway
  }
  full_clone = false
  cloud_init              = true
  cloud_init_storage_pool = "storage01"

  ssh_host             = local.network.ip
  ssh_username         = local.ssh.username
  ssh_private_key_file = local.ssh.key_file
  ssh_timeout          = "10m"
}

build {
  sources = ["source.proxmox-clone.ubuntu"]
  # Install open JRE17
  provisioner "shell" {
    inline = [
      "sudo apt install openjdk-17-jre-headless -y"
    ]
  }
}
