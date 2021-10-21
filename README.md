# Proxmox infrastructure

Document my homelab

This was forked from my [OSX Davinci experiments](https://github.com/daneroo/osx-install-davinci/blob/master/Proxmox.md)

- Docker/K8S for remote development
- Storage
  - `/archive`
  - TimeMachine/CCC Destination
  - Plex and media content

2021-10-21: Proxmox installed on hilbert
2020-08-26: Proxmox Installed on fermat (full disk) with Catalina and Ubuntu-20.04 guest vms.

- [Proxmox on Hilbert](https://hilbert.imetrical.com:8006)

- [Proxmox on Fermat](https://fermat.imetrical.com:8006)
  - fermat.imetrical.com: 192.168.86.239
    - Static IP `192.168.86.239` was assigned in `/etc/network/interfaces`
    - Same IP was added to Google WIFI DHCP Reservation for it (under uneditable name **Apple**)

## TODO

- Install Proxmox Ve 7.0 on hilbert (old DaVinci/Designare)
  - Disk inventory
    - /dev/nvme0n1: 931.5 GiB (Sabrent Rocket 1TB - one of them)
    - /dev/sdb: 2TB - DaVinci20Clone
    - /dev/sdc: 298.1 GiB - old hitachi 2.5" HDD (HiveOS)
  - Move HiveOS to hilbert/vm
  - Update OSX Section for Bug Sur
- Docs: Copy over Benchmarks/Storage docs from `osx-install-davinci/` repo.
- Provisioning with Pulumi
  - [pulumi-proxmox](https://www.npmjs.com/package/@matchlighter/pulumi-proxmoxve)
  - parametrizing cloud-init or building custom iso's and templates
  - adding snaps
    - <https://cloudinit.readthedocs.io/en/latest/topics/format.html>
- Where does Tailscale fit in
  - [Tailscale + k3s](https://blog.dsb.dev/posts/accessing-my-k3s-cluster-from-anywhere-with-tailscale/index.html)
- Integrate all infra into monorepo, combining docs and history and qcic

## Install and Setup

_Took 20 minutes to burn and install!_

Got `proxmox-ve_6.2-1.iso` from <https://www.proxmox.com/en/downloads/category/iso-images-pve>

- Burn iso to DVD worked much better than USB Stick, for booting Mac Mini)
- [Boot Installer (with Alt on non-MacOS keyboard)](https://support.apple.com/en-gb/HT201255)

### Big Sur

- [Installing macOS 11 “Big Sur” on Proxmox 6
  ](https://www.nicksherlock.com/2020/06/installing-macos-big-sur-on-proxmox/)
- [Author's current Setup](https://www.nicksherlock.com/2018/11/my-macos-vm-proxmox-setup/)

### Catalina

- [Bringing up Catalina w/OpenCore](https://www.nicksherlock.com/2020/04/installing-macos-catalina-on-proxmox-with-opencore/)

As in the article above..

```bash
./fetch-macOS.py

hdiutil convert BaseSystem.dmg -format RdWr -o Catalina-installer.iso
mv Catalina-installer.iso.img Catalina-installer.iso

$ ~/Downloads/Proxmox/smc_read
ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc
```

## Making Fresh ubuntu iso for proxmox

It should be possible to install from cloud-init (pointing to a static server with config), or perhaps with the auto_install script baked into the iso/template, but this is not yet done.

When we instantiate a vm from a template, we will still want to provide at a minimum a new hostname. This is what the current `pulumi-pxmx` example does.

- Ubuntu Autoinstall (Subiquity)
  - <https://ubuntu.com/server/docs/install/autoinstall>
  - <https://ubuntu.com/server/docs/install/autoinstall-quickstart>
  - post installation, stored at `/var/log/installer/autoinstall-user-data`
- See <https://www.aerialls.io/posts/ubuntu-server-2004-image-packer-subiquity-for-proxmox/>
- How to make a custom iso (pre-autoinstall a.k.a <18.04) <https://github.com/Telmate/terraform-ubuntu-proxmox-iso>

### Cloud init proxmox template

This is how we made a proxmox template vm, which we can instantiate from pulumi.

```bash
# download the image
wget https://cloud-images.ubuntu.com/releases/focal/release/ubuntu-20.04-server-cloudimg-amd64.img
# create a new VM
qm create 9000 --memory 2048 --net0 virtio,bridge=vmbr0

# import the downloaded disk to local-lvm storage
# qm importdisk 9000 bionic-server-cloudimg-amd64.img local-lvm
qm importdisk 9000 ubuntu-20.04-server-cloudimg-amd64.img local-lvm

# finally attach the new disk to the VM as scsi drive
qm set 9000 --scsihw virtio-scsi-pci --scsi0 local-lvm:vm-9000-disk-1

# Add Cloud-Init CDROM drive : then edit and regenerate
qm set 9000 --ide2 local-lvm:cloudinit

# set the boot disk
qm set 9000 --boot c --bootdisk scsi0

# Also configure a serial console and use it as a display
qm set 9000 --serial0 socket --vga serial0

# convert the VM into a template
qm template 9000
```

### Cloning the template

```bash
#To examine:
qm cloudinit dump 9000 user
qm cloudinit dump 9000 network
qm cloudinit dump 9000 meta

# To Clone
qm clone 9000 123 --name new-name-of-vm
```

## Ubuntu 20.04

- Update boot iso
- check docker at install

```bash
sudo apt install net-tools htop iotop bash-completion -y
# setup docker
sudo groupadd docker
sudo usermod -aG docker $USER
sudo snap restart docker.dockerd
# - or -
sudo snap remove docker
sudo snap install docker

docker run --rm hello-world
## docker bash completion, from this gist
# - https://gist.github.com/toschneck/2df90c66e0f8d4c6567d69a36bfc5bcd

# upgrade
sudo apt update && sudo apt dist-upgrade && sudo apt autoremove
sudo apt-get autoremove --purge

# install go
sudo snap install go --classic # classic may not necessary
go version
sudo apt install build-essential -y # for cgo

# install nodejs - default channel is currently 12/stable
sudo snap install node --channel=14/stable --classic

# install ipfs - not tested - https://snapcraft.io/install/ipfs/ubuntu
sudo snap install ipfs

# setup PATH and bash...
```

### Resize Disk

[Resize disk in Prox mox](https://pve.proxmox.com/wiki/Resize_disks#Online_for_Linux_Guests)

```bash
# Physical PV
pvresize /dev/vda3
# Logical LV
lvresize --extents +100%FREE --resizefs /dev/mapper/ubuntu--vg-ubuntu--lv
```

## References

- [updated Catalina install post (with OpenCore)](https://www.nicksherlock.com/2020/04/installing-macos-catalina-on-proxmox-with-opencore/)
- [Proxmox on Mac Mini](https://wilfredomaldonado.wordpress.com/2017/02/03/proxmox-ve-on-mac-mini-part-1/)
- [Galois setup : dual boot OSX/Ubuntu 20.04](https://www.evernote.com/shard/s60/nl/6925909/3c5f0f1f-8131-459b-b9dd-dc3e9dedda4d/)
