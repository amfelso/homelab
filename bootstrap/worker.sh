CMDLINE_FILE="/boot/firmware/cmdline.txt"

sudo apt update && sudo apt upgrade -y
sudo apt install -y htop nano curl net-tools git

if ! grep -q "cgroup_enable=memory" "$CMDLINE_FILE"; then
    echo "Appending cgroup_enable=memory cgroup_memory=1 to cmdline.txt"
    sudo sed -i 's/$/ cgroup_enable=memory cgroup_memory=1/' "$CMDLINE_FILE"
    echo "Memory cgroups added — rebooting required before installing k3s..."
    sudo reboot
    exit 0
else
    echo "Memory cgroups already enabled, continuing..."
fi

curl -sfL https://get.k3s.io | K3S_URL=$NODE_URL K3S_TOKEN=$NODE_TOKEN sh -s