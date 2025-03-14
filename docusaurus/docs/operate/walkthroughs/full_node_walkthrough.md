---
title: Full Node - Binary (~30 min)
sidebar_position: 1
---

**🧑‍🔬 detailed step-by-step instructions to get you up and running with a `Full Node` on Pocket Network ✅**

:::warning This is an in-depth walkthrough

See the [Full Node Cheat Sheet](../cheat_sheets/full_node_cheatsheet.md) if you want to just copy-pasta a few commands.

**These instructions are intended for Debian-based systems.**

:::

---

## Table of Contents <!-- omit in toc -->

- [Introduction - why run a Full Node?](#introduction---why-run-a-full-node)
- [Pre-Requisites \& Requirements](#pre-requisites--requirements)
- [Instructions](#instructions)
  - [1. Install Dependencies](#1-install-dependencies)
  - [2. \[OPTIONAL\] Create a New User](#2-optional-create-a-new-user)
  - [3. Set Up Environment Variables for Cosmovisor](#3-set-up-environment-variables-for-cosmovisor)
  - [4. Install Cosmovisor](#4-install-cosmovisor)
  - [5. Install `poktrolld`](#5-install-poktrolld)
  - [6. Retrieve the latest genesis file](#6-retrieve-the-latest-genesis-file)
  - [7. Network Configuration](#7-network-configuration)
  - [8. Set Up `systemd` Service](#8-set-up-systemd-service)
  - [9. Configure your Firewall](#9-configure-your-firewall)

## Introduction - why run a Full Node?

This guide will guide through, step-by-step, through running a Full Node for Pocket Network.

Running a Full Node is the first step toward becoming a Validator, Supplier, or Gateway.

The instructions outlined here use [Cosmovisor](https://docs.cosmos.network/v0.45/run-node/cosmovisor.html)
to enable automatic binary upgrades.

## Pre-Requisites & Requirements

1. **Linux-based System**: Preferably Debian-based distributions.
2. **Root or Sudo Access**: Administrative privileges are required.
3. **Dedicated Server or Virtual Machine**: Any provider is acceptable.

## Instructions

### 1. Install Dependencies

1. **Install `curl`, `tar`, `wget`, `jq`**

   ```bash
   sudo apt-get update
   sudo apt-get install -y curl tar wget jq
   ```

2. **Install `go`** by following the official installation instructions [here](https://go.dev/doc/install).

### 2. [OPTIONAL] Create a New User

It is considered good practice to not run the binary as `root`. The below instructions will show you how to create a user with scope to run the node.

Create a dedicated user to run `poktrolld`:

```bash
sudo adduser poktroll
```

Set a password when prompted, and add the user to the sudo group:

```bash
sudo usermod -aG sudo poktroll
```

And switch to the `poktroll` user:

```bash
sudo su - poktroll
```

### 3. Set Up Environment Variables for Cosmovisor

Create a `.poktrollrc` file and set the following environment variables:

```bash
touch ~/.poktrollrc

echo "export DAEMON_NAME=poktrolld" >> ~/.poktrollrc
echo "export DAEMON_HOME=\$HOME/.poktroll" >> ~/.poktrollrc
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.poktrollrc
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=true" >> ~/.poktrollrc
echo "export UNSAFE_SKIP_BACKUP=false" >> ~/.poktrollrc

echo "source ~/.poktrollrc" >> ~/.profile
source ~/.profile
```

### 4. Install Cosmovisor

Follow the official Cosmovisor installation instructions [here](https://docs.cosmos.network/main/build/tooling/cosmovisor#installation).

Add `cosmovisor` to your `$PATH`

```bash
# This uses the default go behavior and may be different if you customize your $GOPATH or $GOBIN
export PATH=$PATH:$HOME/go/bin

# Verify the install
cosmovisor version
```

:::info `cosmovisor` Install Verification

**This MAY output an `Error` message.**

As long as it prints `cosmovisor version: <some_version>`, this worked correctly.

:::

### 5. Install `poktrolld`

Follow the instructions in the [poktrolld Installation Guide](../../tools/user_guide/poktrolld_cli.md) page to install `poktrolld`.

Create a symlink of the binary so Comosvisor knows where to find it:

```bash
mkdir -p $HOME/.poktroll/cosmovisor/genesis/bin
ln -sf $(which poktrolld) $HOME/.poktroll/cosmovisor/genesis/bin/poktrolld
```

### 6. Retrieve the latest genesis file

Follow the instructions below to download the latest genesis file.

```bash
# Select network (testnet-alpha, testnet-beta, or mainnet)
NETWORK="testnet-beta" # Change this to your desired network

# Create config directory if it doesn't exist
mkdir -p $HOME/.poktroll/config

# Download genesis file
GENESIS_URL="https://raw.githubusercontent.com/pokt-network/pocket-network-genesis/master/shannon/${NETWORK}/genesis.json"
curl -s -o $HOME/.poktroll/config/genesis.json "$GENESIS_URL"
```

### 7. Network Configuration

:::note
You may see a message saying `genesis.json file already exists`.

This is expected since we downloaded the genesis file in one of the steps above.

The initialization will still complete successfully and set up the required configuration files.

:::

Run the following commands to configure your network environment appropriately:

```bash
# Extract chain-id from existing genesis
CHAIN_ID=$(jq -r '.chain_id' < $HOME/.poktroll/config/genesis.json)

# Initialize the node
poktrolld init "YourNodeMoniker_REPLACE_ME" --chain-id="$CHAIN_ID" --home=$HOME/.poktroll

# Set the seeds
SEEDS_URL="https://raw.githubusercontent.com/pokt-network/pocket-network-genesis/master/shannon/${NETWORK}/seeds"
SEEDS=$(curl -s "$SEEDS_URL")
sed -i -e "s|^seeds *=.*|seeds = \"$SEEDS\"|" $HOME/.poktroll/config/config.toml

# [OPTIONAL] Set External Address.
# The default external address is the external IP of the host.
EXTERNAL_IP=$(curl -s https://api.ipify.org)
sed -i -e "s|^external_address *=.*|external_address = \"${EXTERNAL_IP}:26656\"|" $HOME/.poktroll/config/config.toml
```

### 8. Set Up `systemd` Service

Create a `systemd` service file to manage the node:

```bash
sudo tee /etc/systemd/system/cosmovisor.service > /dev/null <<EOF
[Unit]
Description=Cosmovisor daemon for poktrolld
After=network-online.target

[Service]
User=poktroll
ExecStart=/home/poktroll/go/bin/cosmovisor run start --home=/home/poktroll/.poktroll
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
Environment="DAEMON_NAME=poktrolld"
Environment="DAEMON_HOME=/home/poktroll/.poktroll"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
Environment="UNSAFE_SKIP_BACKUP=true"

[Install]
WantedBy=multi-user.target
EOF
```

Enable and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable cosmovisor.service
sudo systemctl start cosmovisor.service

# Check the logs of the service
journalctl -u cosmovisor -f
```

### 9. Configure your Firewall

To ensure your node can properly participate in the P2P network, you need to make port `26656` accessible from the internet.

This may involve one or more of the following:

a. **Configuring your firewall for UFW**:

```bash
# [REQUIRED] Expose P2P
sudo ufw allow 26656/tcp

# [OPTIONAL] Expose additional ports
# Enable RPC from Internet (requires config change in .poktroll/config/config.toml)
sudo ufw allow 26657/tcp
```

b. **Configuring your firewall for iptables**:

```bash
sudo iptables -A INPUT -p tcp --dport 26656 -j ACCEPT
```

**Cloud Provider Settings**: If running on a cloud provider (AWS, GCP, Azure, etc.), ensure you configure the security groups or firewall rules to allow inbound traffic on port 26656.
c. **Router Configuration**: If running behind a router, configure port forwarding for port 26656 to your node's internal IP address.
d. **Verify your port** is accessible using a tool like netcat or telnet from another machine:

```bash
nc -vz {EXTERNAL_IP} 26656
```
