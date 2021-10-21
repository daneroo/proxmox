# Talos on Proxmox

- [Reference in Docs](https://www.talos.dev/docs/v0.11/virtualized-platforms/proxmox/)

- [Downloaded iso](https://github.com/talos-systems/talos/releases/tag/v0.11.5)
- Create VM (2core/2GB/10GB)
- Boot to get IP

## TODO

- [Dynamic Providers](https://www.pulumi.com/blog/dynamic-providers/)
## Completing the Installation

- pulumi up
- get mac addresses
- start talos dance

```bash
export CONTROL_PLANE_IP=192.168.86.27
export WORKER_IP=192.168.86.34
export TALOSCONFIGDIR="_out"
export TALOSCONFIG="${TALOSCONFIGDIR}/talosconfig"

./talosctl gen config talos-proxmox https://$CONTROL_PLANE_IP:6443 --output-dir ${TALOSCONFIGDIR}

# This causes a reboot after installation
./talosctl apply-config --insecure --nodes $CONTROL_PLANE_IP --file ${TALOSCONFIGDIR}/controlplane.yaml

# This causes a reboot after installation
./talosctl apply-config --insecure --nodes $WORKER_IP --file ${TALOSCONFIGDIR}/join.yaml

./talosctl config endpoint $CONTROL_PLANE_IP
./talosctl config node $CONTROL_PLANE_IP

./talosctl bootstrap

./talosctl kubeconfig .

kubectl --kubeconfig kubeconfig get nodes

```

## Install talosctl

```bash
curl https://github.com/talos-systems/talos/releases/latest/download/talosctl-darwin-amd64 -L -o talosctl
sudo cp talosctl /usr/local/bin
sudo chmod +x /usr/local/bin/talosctl
```
