# Docker compose
- Là một công cụ giúp định nghĩa và quản lý các ứng dụng Docker với nhiều container dễ dàng hơn.
- Thay vì phải chạy từng container một cách thủ công bằng các lệnh Docker riêng lẻ, Docker Compose cho phép định nghĩa tất cả container, các cấu hình của chúng, và cách chúng kết nối với nhau trong một file duy nhất, thường là `docker-compose.yml.`

## Các thành phần chính của Docker Compose:
- **Service (Dịch vụ)**: Mỗi service tương ứng với một container. Có thể định nghĩa nhiều service trong một ứng dụng, như một service chạy cơ sở dữ liệu (MySQL), một service chạy ứng dụng web (Node.js), v.v.
- **Network (Mạng)**: Docker Compose tự động tạo ra một mạng riêng cho các service để chúng có thể giao tiếp với nhau.
- **Volume (Ổ đĩa ảo)**: Là cách để lưu trữ dữ liệu lâu dài, không bị mất khi container bị xoá hoặc dừng.
- **Dependencies (Phụ thuộc)**: Docker Compose có thể định nghĩa thứ tự khởi động giữa các service, ví dụ: container ứng dụng web chỉ được khởi động sau khi container cơ sở dữ liệu đã sẵn sàng.

## Ví dụ Dockercompose cơ bản

```
version: '3'
services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
  db:
    image: mysql:latest
    environment:
      MYSQL_ROOT_PASSWORD: example
```

- Service "web" chạy container Nginx, và ánh xạ cổng 8080 của máy chủ đến cổng 80 của container.
- Service "db" chạy container MySQL với biến môi trường để đặt mật khẩu gốc cho MySQL.
- Khi chạy lệnh docker-compose up, cả hai container này sẽ được khởi động và kết nối với nhau theo cấu hình được định nghĩa.
