# ubuntu host for docker workloads

Created VM on px1 with:

- 4 GB / 4 cores / 100GB disk
- ubuntu-22.04.1-live-server-amd64.iso

## Ubuntu install 22.04 Server

- update to new installer: yes
- Base: Ubuntu Server (no minimized)
- ip: dhcp: 192.168.86.240/24
- Snaps: select none (docker by apt)
- detach cd-rom and reboot

## Post-install

- [x] Hover DNS hostname: d1-px1.imetrical.com : 192.168.86.240
- [x] Reserved on Google Wifi
- [ ] fqdn: d1-px1.imetrical.com in `/etc/hosts`
  - `127.0.1.1 d1-px1.imetrical.com d1-px1`
- [x] [install tailscale](https://tailscale.com/kb/1015/install/)
- [x] tailscale
- [x] [install docker (with apt)](https://docs.docker.com/engine/install/ubuntu/)
- [x] docker permissions

```bash
sudo groupadd docker
sudo usermod -aG docker $USER
sudo snap restart docker.dockerd
```

## deploy scrobblecast

- add ssh key to github

```bash
# checkout scrobblecast
mkdir -p Code/iMetrical
cd Code/iMetrical/
git clone git@github.com:daneroo/scrobbleCast.git

# Copy credentials from dirac
daniel@dirac$ scp -p credentials.* s3cfg.env daniel@d1-px1:Code/iMetrical/scrobbleCast/js
# Copy data from darwin
daniel@darwin$ scp -rp data daniel@d1-px1.imetrical.com:Code/iMetrical/scrobbleCast/js
daniel@darwin$ rsync -av -n data/ daniel@d1-px1.imetrical.com:Code/iMetrical/scrobbleCast/js/data/


```

- [ ] copy seed data
- [ ] make start
