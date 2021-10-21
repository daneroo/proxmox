import * as pulumi from '@pulumi/pulumi'
import * as pve from '@matchlighter/pulumi-proxmoxve'

const qemuConfig = {
  // name: 'name of this vm',
  // desc: 'description of this vm',
  targetNode: 'fermat',
  iso: 'local:iso/talos-amd64.iso',

  memory: 2048,
  sockets: 1,
  cores: 2,
  scsihw: 'virtio-scsi-pci', // default would not produce /dev/sda
  // ipconfig0: 'ip=dhcp', // only for cloudInit?
  networks: [
    {
      model: 'virtio',
      // firewall: true, // Whether to enable the Proxmox firewall on this network device
      bridge: 'vmbr0'
    }
  ],
  disks: [
    {
      storage: 'local-lvm',
      type: 'scsi',
      size: '10G'
    }
  ]
}
const talos0 = new pve.QemuVM('talos-0', qemuConfig)
const talos1 = new pve.QemuVM('talos-1', qemuConfig)

export const talos0vmid = talos0.vmid
export const talos1vmid = talos1.vmid
// export const macaddr = talos0.networks[0].macaddr
// export const json = JSON.stringify({ talos0, talos1 })
// export const talos0macaddr = talos0.networks[0].macaddr
// export const talos1macaddr = talos1.networks[0].macaddr

// export const t0json = pulumi.output(talos0).apply(talos => {
//   // t.networks[0].macaddr
//   return JSON.stringify(talos)
// })
// pulumi.log.info('t0json', t0json)
