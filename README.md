# Homelab

[![.github/workflows/develop.yaml](https://github.com/amfelso/homelab/actions/workflows/develop.yaml/badge.svg)](https://github.com/amfelso/homelab/actions/workflows/develop.yaml)
[![.github/workflows/main.yaml](https://github.com/amfelso/homelab/actions/workflows/main.yaml/badge.svg?branch=main)](https://github.com/amfelso/homelab/actions/workflows/main.yaml)

This repo contains automation scripts, application code, and Kubernetes manifests for a Raspberry Pi K3s homelab.

## Prerequisites

- **Raspberry Pi nodes:** Minimum 2 (4× Model B recommended, ≥2 GB RAM, 32 GB microSD).
- **Operating system:** Flash each microSD with Ubuntu Server ≥25.10 (64-bit).
- **Admin user & SSH:** Create an admin user with SSH access, and store the username and private key path in `.env`.
- **Network:** Connect Pis via Ethernet and assign stable IPs (e.g., via DHCP reservation); store them in `.env`.
- **Cluster role:** `NODE_01` will be the K3s master; other nodes will be workers.

## Quickstart

1. **Bootstrap the master node:**

```bash
make bootstrap node=1
```

On the first run, the Pi may reboot after enabling memory cgroups. Wait for it to come back online and rerun the command.
After bootstrapping the K3s master, your cluster config and node token will be available for worker nodes.

2. **Bootstrap worker node(s):**

```bash
make bootstrap node=<node-ip>
```

3. **Activate the virtual environment to run `kubectl` commands:**

```bash
source venv/activate
```
