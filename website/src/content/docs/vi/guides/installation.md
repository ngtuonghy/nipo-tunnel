---
title: Cài đặt
description: Hướng dẫn cài đặt và gỡ cài đặt Nipo.
sidebar:
  order: 2
---

Hãy làm theo các bước sau để cài đặt và sử dụng Nipo.

## Yêu cầu hệ thống

* **Node.js** (v14 trở lên) - được dùng để tải công cụ dòng lệnh toàn cục.

## Cài đặt

Cài đặt Nipo trên toàn hệ thống bằng npm:

```bash
npm install -g nipo-tunnel
```

## Gỡ cài đặt

Để xóa hoàn toàn Nipo Tunnel và mọi cấu hình của nó khỏi hệ thống của bạn:

```bash
# 1. Gỡ cài đặt gói npm
npm uninstall -g nipo-tunnel

# 2. Xóa các cấu hình và file lõi
# Trên Windows (PowerShell)
Remove-Item -Recurse -Force ~/.nipo
# Trên Mac/Linux
rm -rf ~/.nipo
```
