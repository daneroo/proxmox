# Proxmox test installation

This was forked from my [OSX Davinci experiments](https://github.com/daneroo/osx-install-davinci/blob/master/Proxmox.md)

Objective: manage cluster (k8s) and remote workspaces

- Docker remote for development
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

# setup PATH and bash...
```

```bash
# session
  136  find /usr/local/
  137  export PATH=/usr/local/go/bin:$PATH
  138  go vewrsion
  139  go version
  140  wget -c https://dl.google.com/go/go1.15.5.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
  141  go version
  142  time go run scripts/pump.go
  143  docker ps -a
  144  time go run scripts/pump.go
  145  go mod tidy
  146  ls
  147  rm go1.15.5.linux-amd64.tar.gz
  148  rm go1.15.5.linux-amd64.tar.gz.1
  149  git log
  150  git checkout develop
  151  git status
  152  go get gopls
  153  go mod tidy
  154  time go run scripts/pump.go
  155  time scp -p daniel@dirac.imetrical.com:Code/Go/src/github.com/daneroo/go-ted1k/data .
  156  time scp -rp daniel@dirac.imetrical.com:Code/Go/src/github.com/daneroo/go-ted1k/data .
  157  history
```

## References

- [updated Catalina install post (with OpenCore)](https://www.nicksherlock.com/2020/04/installing-macos-catalina-on-proxmox-with-opencore/)
- [Proxmox on Mac Mini](https://wilfredomaldonado.wordpress.com/2017/02/03/proxmox-ve-on-mac-mini-part-1/)
- [Galois setup : dual boot OSX/Ubuntu 20.04](https://www.evernote.com/shard/s60/nl/6925909/3c5f0f1f-8131-459b-b9dd-dc3e9dedda4d/)
