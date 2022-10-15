package provider

// Proxmox https://www.packer.io/plugins/builders/proxmox/iso
type Proxmox struct {
	// Required https://www.packer.io/plugins/builders/proxmox/iso#required
	proxmoxUrl string
	// Username when authenticating to Proxmox, including the realm.
	username string
	// Either password or token
	password string
	// Either password or token
	token string
	// Path to the ISO file to boot from, expressed as a proxmox datastore path iso_file OR iso_url must be specifed.
	isoFile string
	// URL to an ISO file to upload to Proxmox, and then boot from. Either iso_file OR iso_url must be specifed.
	isoURL string
	// Proxmox storage pool onto which to upload the ISO file.
	isoStoragePool string
	//isoChecksum    string
	// END Required

	//  If true, remove the mounted ISO from the template after finishing. Defaults to false
	unmountIso bool
	// Skip validating the certificate.
	insecureSkipTlsVerify bool
	// Enables QEMU Agent option for this VM. When enabled, then qemu-guest-agent must be installed on the guest.
	// When disabled, then ssh_host should be used. Defaults to true.
	qemuAgent bool
	//Defaults to 1 minute.
	//taskTimeout time.Duration
}

type System struct {
	// If not given, the next free ID on the node will be used.
	vmID uint16
	// If not given, a random uuid will be used.
	vmName string
	// Defaults to 512
	memory uint8
	// How much memory, in megabytes, to give the virtual machine. Defaults to 512
	cores uint8
	// How many CPU sockets to give the virtual machine. Defaults to 1
	sockets uint8

}

type Disk struct {
	// The size of the disk, including a unit suffix, such as 10G to indicate 10 gigabytes.
	diskSize string
	// How to cache operations to the disk. Can be none, writethrough, writeback, unsafe or directsync. Defaults to none.
	cacheMode string
	// The format of the file backing the disk. Can be raw, cow, qcow, qed, qcow2, vmdk or cloop. Defaults to raw.
	format string
	// Required. Name of the Proxmox storage pool to store the virtual machine disk on. A local-lvm pool is allocated by the installer, for example.
	storagePool string
	// Required. The type of the pool, can be lvm, lvm-thin, zfspool, cephfs, rbd or directory.
	storagePoolType string
	// The type of disk. Can be scsi, sata, virtio or ide. Defaults to scsi.
	type string
	// The SCSI controller model to emulate. Can be lsi, lsi53c810, virtio-scsi-pci, virtio-scsi-single, megasas, or pvscsi. Defaults to lsi.
	scsiController string
}

func New() {

}