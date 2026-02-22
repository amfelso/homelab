#!/bin/bash
set -euo pipefail

# Get a node's IP given node ID
NODE_ID=$1
NODE_NAME_PREFIX="NODE_0"
[ -z "$NODE_ID" ] && { echo "Node ID not set."; exit 1; }
source venv/activate > /dev/null
NODE_VAR="$NODE_NAME_PREFIX$NODE_ID"
echo "${!NODE_VAR}"