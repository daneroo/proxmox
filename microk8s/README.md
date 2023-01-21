# Microk8s on Proxmox

- provision ubuntu-20.04.03 vm
- install microk8s
- configure microk8s
  - addons
  - tailscale?
- use from external workstation

## Provisioning

### Manually

Let's do this manually until pulumi can do the whole flow.

- upload the latest `ubuntu-20.04.3-live-server-amd64.iso` image to the proxmox server
- provision machine name microk8s / 100G/2x2cores/4GB
  - openssh: yes, import keys from github: not yet
  - microk8s: yes (stable channel)
  - remove cd-rom iso, reboot
  - Pin the ip address
    - pin DHCP address (Google Home only)
    - add hover ip address [optional]
  - microk8s enable storage -  which has:
    - name=microk8s-hostpath
    - annotation: storageclass.kubernetes.io/is-default-class=true

```bash
usermod -a -G microk8s daniel
chown -f -R daniel ~/.kube

microk8s status --wait-ready
# enable addons dns storage metrics traefik  metalb openebs
microk8s enable dns storage

# get the certificate
microk8s config
```

### Provisioning from template

- move docs from `../README.md`
- use pulumi to create a new vm from the template

## First app

We will use grafana + postgresql + go-ted1k as our first example

This is not very good, but a start:

- [postgresql](https://severalnines.com/database-blog/using-kubernetes-deploy-postgresql)
- [with helm](https://itnext.io/basic-postgres-database-in-kubernetes-23c7834d91ef)
- [crunchy data PGO operator](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/)
- [helm + openebs](https://thenewstack.io/tutorial-deploy-postgresql-on-kubernetes-running-the-openebs-storage-engine/)