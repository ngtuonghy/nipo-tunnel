---
title: Installation
description: How to install and uninstall Nipo.
sidebar:
  order: 2
---

Follow these instructions to install and use Nipo.

## Prerequisites

* **Node.js** (v14 or higher) - used to install the global binary runner.

## Installation

Install Nipo globally using npm:

```bash
npm install -g nipo-tunnel
```

## Uninstallation

To completely remove Nipo Tunnel and its configuration from your system:

```bash
# 1. Uninstall the npm package
npm uninstall -g nipo-tunnel

# 2. Remove configuration and core binary
# Windows (PowerShell)
Remove-Item -Recurse -Force ~/.nipo
# Mac/Linux
rm -rf ~/.nipo
```
