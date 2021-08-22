
# Pulumi and ProxMox

- <https://github.com/matchlighter/pulumi-proxmoxve>
  - https://github.com/Telmate/proxmox-api-go


## Moving to Talos

## Last State 2021-05-09

Destroyed the stack, but did not remove it!

```bash
pulumi up

pulumi destroy -s daneroo/pulumi-pxmx/dev
# but not removed
# pulumi stack rm daneroo/pulumi-pxmx/dev
```

## Setup from MacOS (darwin)

```bash
wget https://github.com/Matchlighter/pulumi-proxmoxve/releases/download/v0.1.5/pulumi-resource-proxmoxve-v0.1.5-darwin-amd64.tar.gz

pulumi plugin ls

pulumi plugin install resource proxmoxve 0.1.5 -f pulumi-resource-proxmoxve-v0.1.5-darwin-amd64.tar.gz

pulumi plugin ls

```

## pulumi project

```bash
pulumi new # typescript
npm i ts-standard -D
npm install @matchlighter/pulumi-proxmoxve

# user is probably root@pam
pulumi config set --secret proxmoxve:pmUser root@pam
pulumi config set --secret proxmoxve:pmPassword PASSWORD
pulumi config set proxmoxve:pmTlsInsecure true
pulumi config set proxmoxve:pmApiUrl https://fermat.imetrical.com:8006/api2/json
```
