
# Pulumi and ProxMox

- [Dynamic Providers](https://www.pulumi.com/blog/dynamic-providers/)
- <https://github.com/matchlighter/pulumi-proxmoxve>
  - https://github.com/Telmate/proxmox-api-go

## Moving to Talos

Replacing the ubuntu iso from before with the Talos iso.

```bash
pulumi up
pulumi stack export | jq -r .deployment.resources[2].outputs.networks[0].macaddr
pulumi stack export | jq -r .deployment.resources[3].outputs.networks[0].macaddr
ping 192.168.86.255
# apt-get install net-tools
arp -a |grep macaddress

# nmap
#  ping scan
nmap -sn 192.168.86.0/24

nmap -sP 192.168.86.0/24 >/dev/null && arp -an | grep <mac address here> | awk '{print $2}' | sed 's/[()]//g'

arp -a
? (192.168.86.25) at 82:11:d0:34:89:b on en0 ifscope [ethernet]
? (192.168.86.27) at (incomplete) on en0 ifscope [ethernet]
? (192.168.86.28) at 4a:98:5a:b5:d9:9a on en0 ifscope [ethernet]
```

## nmap in Go

- [nmap Golang](https://github.com/Ullaakut/nmap)

```bash
git clone git@github.com:Ullaakut/nmap.git
```

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
