# Proxmox test installation

THis was forked from my [OSX Davinci experiments](https://github.com/daneroo/osx-install-davinci/blob/master/Proxmox.md)

## Install

See: <https://pve.proxmox.com/wiki/Install_from_USB_Stick>

### USB Stick

Got `proxmox-ve_6.1-1.iso` from <https://www.proxmox.com/en/downloads/category/iso-images-pve>

*or try with `balenaEtcher`*

```bash
hdiutil convert -format UDRW -o proxmox-ve_6.1-1.dmg proxmox-ve_6.1-1.iso

diskutil list # DataTraveller is disk6
diskutil unmountDisk /dev/disk6
sudo dd if=proxmox-ve_6.1-1.dmg of=/dev/rdisk6 bs=1m
```

### Catalina

See: <https://www.nicksherlock.com/2019/10/installing-macos-catalina-10-15-on-proxmox-6/>

```bash
$ ~/Downloads/Proxmox/smc_read
ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc
```
