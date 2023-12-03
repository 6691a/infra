packer {
  required_plugins {
    proxmox = {
      version = " >= 1.1.5"
      source  = "github.com/hashicorp/proxmox"
    }
  }
}

source "proxmox-iso" "ubuntu" {
  proxmox_url              = local.credentials.api_url
  username                 = local.credentials.api_token_id
  token                    = local.credentials.api_token_secret
  insecure_skip_tls_verify = true

  node    = local.vm.node
  vm_id   = local.vm.id
  vm_name = local.vm.name

  iso_file         = local.vm.iso
  iso_storage_pool = local.vm.iso_storage_pool
  unmount_iso      = true
  qemu_agent = true

  boot_command = [
    "<esc><wait>",
    "e<wait>",
    "<down><down><down><end>",
    "<bs><bs><bs><bs><wait>",
    "autoinstall ds=nocloud-net\\;s=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ ---<wait>",
    "<f10><wait>"
  ]
  boot = "c"
  boot_wait = "10s"

  scsi_controller = "virtio-scsi-pci"
  disks {
    disk_size    = local.disk.size
    storage_pool = local.disk.storage_pool
    type         = local.disk.type
    format       = local.disk.format
  }

  cores  = local.vm.cores
  memory = local.vm.memory

  network_adapters {
    bridge = "vmbr0"
    model  = "virtio"
    firewall = "false"
  }

  http_directory = "http"
  cloud_init              = true
  cloud_init_storage_pool = "storage01"

  ssh_timeout          = local.ssh.timeout
  ssh_username         = local.ssh.username
  ssh_private_key_file = local.ssh.key_file
}

build {
  sources = ["source.proxmox-iso.ubuntu"]

  provisioner "shell" {
    inline = [
        "while [ ! -f /var/lib/cloud/instance/boot-finished ]; do echo 'Waiting for cloud-init...'; sleep 1; done",
        "sudo rm /etc/ssh/ssh_host_*",
        "sudo truncate -s 0 /etc/machine-id",
        "sudo apt -y autoremove --purge",
        "sudo apt -y clean",
        "sudo apt -y autoclean",
        "sudo cloud-init clean",
        "sudo rm -f /etc/cloud/cloud.cfg.d/subiquity-disable-cloudinit-networking.cfg",
        "sudo rm -f /etc/netplan/00-installer-config.yaml",
        "sudo sync"
    ]
  }

    # Provisioning the VM Template for Cloud-Init Integration in Proxmox #2
  provisioner "file" {
      source = "files/99-pve.cfg"
      destination = "/tmp/99-pve.cfg"
  }

    # Provisioning the VM Template for Cloud-Init Integration in Proxmox #3
  provisioner "shell" {
      inline = [ "sudo cp /tmp/99-pve.cfg /etc/cloud/cloud.cfg.d/99-pve.cfg" ]
  }
}
