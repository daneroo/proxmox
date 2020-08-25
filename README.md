# Proxmox test installation

This was forked from my [OSX Davinci experiments](https://github.com/daneroo/osx-install-davinci/blob/master/Proxmox.md)

Objective: manage cluster (k8s) and remote workspaces

- Storage
- Screen Sharing - CHromebook/MacOS

## TODO

- Could try on fermat. (Core i7?)
- [Catalina w/OpenCore](https://www.nicksherlock.com/2020/04/installing-macos-catalina-on-proxmox-with-opencore/)
- [Galois setup : dual boot OSX/Ubuntu 20.04](https://www.evernote.com/shard/s60/nl/6925909/3c5f0f1f-8131-459b-b9dd-dc3e9dedda4d/)

## Install

See: <https://pve.proxmox.com/wiki/Install_from_USB_Stick>

Boot Installer (with Alt on non MacOS keyboard)

### USB Stick

Got `proxmox-ve_6.2-1.iso` from <https://www.proxmox.com/en/downloads/category/iso-images-pve>

- Burn to USB with [Balena Etcher](https://www.balena.io/etcher/)

### Catalina

See [updated Catalina install post (with OpenCore)](https://www.nicksherlock.com/2020/04/installing-macos-catalina-on-proxmox-with-opencore/)

```bash
$ ~/Downloads/Proxmox/smc_read
ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc
```
