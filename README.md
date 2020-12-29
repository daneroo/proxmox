# Proxmox test installation

This was forked from my [OSX Davinci experiments](https://github.com/daneroo/osx-install-davinci/blob/master/Proxmox.md)

Objective: manage cluster (k8s) and remote workspaces

- Docker remote for development
- ipfs
- Storage
- Screen Sharing - Chromebook/MacOS
- Storage redundancy
- Kubera on K8S on Proxmox

State 2020-08-26: Proxmox Installed on fermat (full disk) with Catalina and Ubuntu-20.04 guest vms.

- [Proxmox on Fermat](https://fermat.imetrical.com:8006)
- Hover: DNS Name fermat.imetrical.com: 192.168.86.239
- Static IP `192.168.86.239` was assign in `/etc/network/interfaces`
- Same IP was added to Google WIFI DHCP Reservation for it (under uneditable name **Apple**)

## TODO

- Where does Tailscale fit in
- Integrate with master infra repo

## Install and Setup

_Took 20 minutes to burn and install!_

Got `proxmox-ve_6.2-1.iso` from <https://www.proxmox.com/en/downloads/category/iso-images-pve>

- Burn iso to DVD worked much better than USB Stick, for booting Mac Mini)
- [Boot Installer (with Alt on non-MacOS keyboard)](https://support.apple.com/en-gb/HT201255)

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
