# Experiment with Talos k8s over Sidero Labs Omni

Omni is a hosted control plane for bare metal clusters. It is a SaaS product.

- [My hosted endpoint](https://imetrical.omni.siderolabs.io/omni/)
- [Omni Docs](https://omni.siderolabs.com/docs/)

## Setup

- Download custom image for omni (e.g. `talos-amd64-v1.2.7-omni-imetrical.iso`)
- Upload image to proxmox

### Create VMs

- Create a VM with 4 cores, 4GB RAM, 32GB disk using this iso
- **Note: As of Talos v1.0 (which requires the x86-64-v2 microarchitecture), booting with the default Processor Type kvm64 will not work.**
  - [Edit the proxmox vm config](https://www.talos.dev/v1.3/talos-guides/install/virtualized-platforms/proxmox/)
  - `/etc/pve/qemu-server/<vmid>.conf`
  - add `args: -cpu kvm64,+cx16,+lahf_lm,+popcnt,+sse3,+ssse3,+sse4.1,+sse4.2`
- Clone it to make 3 machines, total

### Create cluster

- Go to [Omni console](https://imetrical.omni.siderolabs.io/omni/)

### Connect to cluster

Download kubeconfig, e.g. `talos-default-kubeconfig.yml`, and set `KUBECONFIG`.
This requires the [kubelogin kubectl extension](https://github.com/int128/kubelogin) for oicd auth.
Or `brew install int128/kubelogin/kubelogin` on macOS.

```bash
export KUBECONFIG=~/Downloads/talos-default-kubeconfig.yml 
kubectl get nodes
kubectl get nodes,pods,deployment,svc -o wide

# or
kubectl --kubeconfig talos-default-kubeconfig.yml  get nodes
```

### Simple pod deployment

- <https://github.com/kubernetes/examples/blob/master/staging/simple-nginx.md>
- <https://github.com/stefanprodan/podinfo>
- expose with ClusterIp

```bash
kubectl get nodes,pods,deployment,svc -o wide

# --image stefanprodan/podinfo
kubectl create deployment podinfo --image=stefanprodan/podinfo:6.3.0
kubectl scale deployment --replicas 2 podinfo

# implicitly --type=ClusterIP
kubectl expose deployment podinfo --port=9898 

kubectl set image deployment podinfo podinfo=stefanprodan/podinfo:6.2.3

kubectl port-forward svc/podinfo 9898
watch -n 0.5 curl --silent http://localhost:9898
hey http://localhost:9898
stern podinfo

kubectl delete svc,deploy podinfo

# kubectl exec -it pod/podinfo-5d7644b4bb-xpnzk -- ash
```
