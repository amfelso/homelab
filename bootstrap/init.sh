#!/bin/bash

# Load environment, set globals
set +a && source /tmp/.env
LOCAL_IP=$(hostname -I | awk '{print $1}')
CMDLINE_FILE="/boot/firmware/cmdline.txt"

# Update packages, install requirements
sudo apt update && sudo apt upgrade -y
sudo apt install -y htop nano curl net-tools git

# Enable memory cgroups
if ! grep -q "cgroup_enable=memory" "$CMDLINE_FILE"; then
    echo "Appending cgroup_enable=memory cgroup_memory=1 to cmdline.txt"
    sudo sed -i 's/$/ cgroup_enable=memory cgroup_memory=1/' "$CMDLINE_FILE"
    echo "Memory cgroups added — rebooting required before installing k3s..."
    sudo reboot
    exit 0
else
    echo "Memory cgroups already enabled, continuing..."
fi

#Install k3s
if [ "$LOCAL_IP" = "$NODE_01" ]; then
    curl -sfL https://get.k3s.io | sh -

    # Save token and kubeconfig for export
    sed -i "/^NODE_TOKEN=/d" /tmp/.env && \
        echo "NODE_TOKEN=$(sudo cat /var/lib/rancher/k3s/server/node-token)" >> /tmp/.env
    sudo cat /etc/rancher/k3s/k3s.yaml | \
        sed "s/127.0.0.1/$LOCAL_IP/" > /tmp/cluster.yaml

else
    curl -sfL https://get.k3s.io | K3S_URL=$NODE_URL K3S_TOKEN=$NODE_TOKEN sh -s
fi