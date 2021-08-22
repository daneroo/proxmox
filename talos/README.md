# Talos on Proxmox

- [Reference in Docs](https://www.talos.dev/docs/v0.11/virtualized-platforms/proxmox/)

- [Downloaded iso](https://github.com/talos-systems/talos/releases/tag/v0.11.5)
- Create VM (2core/2GB/10GB)
- Boot to get IP

## Thing

```bash
export CONTROL_PLANE_IP=192.168.86.32
./talosctl gen config talos-proxmox https://$CONTROL_PLANE_IP:6443 --output-dir _out
./talosctl apply-config --insecure --nodes $CONTROL_PLANE_IP --file _out/controlplane.yaml

export WORKER_IP=192.168.86.33
./talosctl apply-config --insecure --nodes $WORKER_IP --file _out/join.yaml

export TALOSCONFIG="_out/talosconfig"
./talosctl config endpoint $CONTROL_PLANE_IP
./talosctl config node $CONTROL_PLANE_IP

./talosctl --talosconfig $TALOSCONFIG config endpoint $CONTROL_PLANE_IP
./talosctl --talosconfig $TALOSCONFIG config node $CONTROL_PLANE_IP

./talosctl --talosconfig $TALOSCONFIG bootstrap

./talosctl --talosconfig $TALOSCONFIG kubeconfig .

kubectl --kubeconfig kubeconfig get nodes

```

## Install talosctl

```bash
curl https://github.com/talos-systems/talos/releases/latest/download/talosctl-darwin-amd64 -L -o talosctl
sudo cp talosctl /usr/local/bin
sudo chmod +x /usr/local/bin/talosctl
```
