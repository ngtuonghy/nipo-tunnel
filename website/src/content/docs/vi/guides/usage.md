---
title: Cách sử dụng
description: Hướng dẫn sử dụng Nipo.
sidebar:
  order: 3
---

## 1. Mở cổng HTTP (HTTP Tunneling)
Để điều hướng lưu lượng truy cập công cộng vào một cổng cục bộ (ví dụ: port 3000):
```bash
nipo http 3000
```

## 2. Subdomain tùy chỉnh
Khởi chạy đường hầm với một subdomain (tên miền phụ) tùy chỉnh:
```bash
nipo http 3000 --subdomain myapp
# Hoặc viết tắt:
nipo http 3000 -s myapp
```

## 3. Chạy nhiều đường hầm (thông qua nipo.yml)
Bạn có thể định nghĩa nhiều đường hầm trong cùng một file cấu hình tên là `nipo.yml`:
```yaml
default_subdomain: "myproject"
tunnels:
  - name: web
    port: 3000                     # Sử dụng subdomain mặc định: myproject-web
  - name: api
    port: 8080
    subdomain: "my-custom-api"  # Ghi đè và sử dụng subdomain: my-custom-api
```
* **Khởi chạy toàn bộ đường hầm cùng lúc**:
  ```bash
  nipo start
  ```
* **Khởi chạy một đường hầm cụ thể trong cấu hình**:
  ```bash
  nipo start web
  ```

## 4. Cấu hình & Tùy chọn ngôn ngữ
Kiểm tra các thiết lập cấu hình hiện tại:
```bash
nipo config
```
Nipo tự động nhận diện ngôn ngữ máy tính của bạn, nhưng bạn có thể thay đổi ngôn ngữ hiển thị bất cứ lúc nào:
```bash
nipo config --language vi  # Chuyển sang Tiếng Việt
nipo config --language en  # Chuyển sang Tiếng Anh
```
Những tùy chọn này sẽ được lưu toàn cục trong `~/.nipo/config.yml`.
