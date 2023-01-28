# Re-install Proxmox v7 onto px1.imetrical.com

This is the initial install of Proxmox v7 onto `px1` nee `fermat`.

## Install Proxmox v7

- Boot from USB - Etcher'd `proxmox-ve_7.3-1.iso`
  - monitor and keyboard *directly* attached to mac mini
  - ALT key (on PC Keyboard) pressed to invoke Startup Manager boot from USB
  - Select EFI: Yellow USB icon
  - Install to 1.86TiB Micron SSD as ext4
- hostname: `px1.imetrical.com`
- ip: `192.168.86.239`
- reserve DHCP IP in Google Wifi (update name)
- map name in Hover DNS

## Post-install

- Disable Enterprise subscription repo
  - `cd /etc/apt/sources.list.d`
  - `mv pve-enterprise.list pve-enterprise.list.disabled`
- Add `No-Subscription` repository in `Updates > Repositories`
  - which adds the following to `/etc/apt/sources.list`
  - `deb http://download.proxmox.com/debian/pve bullseye pve-no-subscription`

## Previous settings (v6)

See root's command history before re-install on `./fermat-v6-history.txt`.

- It has traces of Catalina (mostly failed experiment).
- It has kali and min VM's
- It has traces of a template VM (9000) for pulumi-pxmx cloudinit experiment (pre talos).
- It has traces of nested virtualization for my 2022-08-03-PXMX-Clone experiment.
See `proxmox-hilbert/Experiments/2022-08-03-PXMX-Clonne.md`.

```bash
echo "options kvm-intel nested=Y" > /etc/modprobe.d/kvm-intel.conf
modprobe -r kvm_intel
modprobe kvm_intel
modprobe -r kvm_intel
egrep '(vmx|svm)' --color=always /proc/cpuinfo
```
