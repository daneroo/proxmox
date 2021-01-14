import * as pulumi from '@pulumi/pulumi'
import * as pve from '@matchlighter/pulumi-proxmoxve'

const config = new pulumi.Config()
// Can;t get the key from 'proxmoxve:' namespace
// const pmUser = config.requireSecret('proxmoxve:pmUser')
// const pmPassword = config.requireSecret('proxmoxve:pmPassword')
// console.log(`Using auth: ${pmUser}/${pmPassword}`)

const px1vm1 = new pve.QemuVM('px1vm1', {
  targetNode: 'fermat',
  clone: 'VM 9000',
  ciuser: 'daniel',
  cipassword: 'sekret',
  // fullClone: false,
  memory: 2048,
  sockets: 1,
  cores: 2,
  ipconfig0: 'ip=dhcp',
  networks: [{
    model: 'virtio',
    bridge: 'vmbr0'
    // tag: 10
  }],
  disks: [
    {
      storage: 'local-lvm',
      type: 'scsi',
      size: '20G'
    }
  ]
})
