---
title: Usage
description: How to use Nipo.
sidebar:
  order: 3
---

## 1. HTTP Tunneling
To forward public traffic to a local port (e.g., port 3000):
```bash
nipo http 3000
```

## 2. Custom Subdomains
Start a tunnel with a custom subdomain:
```bash
nipo http 3000 --subdomain myapp
# Or shorthand:
nipo http 3000 -s myapp
```

## 3. Multi-tunnel Execution (via nipo.yml)
You can define multiple tunnels in a local configuration file named `nipo.yml`:
```yaml
default_subdomain: "myproject"
tunnels:
  - name: web
    port: 3000                     # Uses default subdomain: myproject-web
  - name: api
    port: 8080
    subdomain: "my-custom-api"  # Overrides and uses: my-custom-api
```
* **Start all tunnels simultaneously**:
  ```bash
  nipo start
  ```
* **Start specific tunnels from the config file**:
  ```bash
  nipo start web
  ```

## 4. Configurations & Language Preferences
Check the active configuration settings:
```bash
nipo config
```
Nipo detects your system locale, but you can change the display language dynamically:
```bash
nipo config --language vi  # Switch to Vietnamese
nipo config --language en  # Switch to English
```
These preferences are saved globally in `~/.nipo/config.yml`.
