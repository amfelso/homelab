# Homelab

This repo contains scripts, configs, and environment files for my Raspberry Pi K3s homelab.

## Quickstart

1. Load environment variables:

```bash
source .env
```

2. SSH into a Pi node:

```bash
ssh -i "$SSH_KEY" $USER@$NODE_MASTER
```
