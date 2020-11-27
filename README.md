# Proxmox test installation

This was forked from my [OSX Davinci experiments](https://github.com/daneroo/osx-install-davinci/blob/master/Proxmox.md)

Objective: manage cluster (k8s) and remote workspaces

- Storage
- Screen Sharing - Chromebook/MacOS
- Storage redundancy
- Kubera on K8S on Proxmox

State 2020-08-26: Proxmox Installed on fermat (full disk) with Catalina and Ubuntu-20.04 guest vms.

- [Proxmox on Fermat](https://192.168.86.239:8006>

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

## References

- [updated Catalina install post (with OpenCore)](https://www.nicksherlock.com/2020/04/installing-macos-catalina-on-proxmox-with-opencore/)
- [Proxmox on Mac Mini](https://wilfredomaldonado.wordpress.com/2017/02/03/proxmox-ve-on-mac-mini-part-1/)
- [Galois setup : dual boot OSX/Ubuntu 20.04](https://www.evernote.com/shard/s60/nl/6925909/3c5f0f1f-8131-459b-b9dd-dc3e9dedda4d/)
