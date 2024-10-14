# Docker images
## Image trong Docker là các phần mềm được đóng gói và quản lý bởi Docker
- Một image là mẫu (template) của container. Nó chứa mọi thứ mà ứng dụng cần để chạy, ví dụ: mã nguồn, công cụ, môi trường...
- Có thể tưởng tượng image như một chiếc khuôn mẫu, từ đó tạo ra nhiều container giống nhau.

**Các lệnh cơ bản**
- Liệt kê image:
    ```
    docker image
    ```

- Tải image:
    ```
  docker pull /image:tag/
    ```

- Xem các container đang chạy:

    `docker ps`

- Truy cập vào container:

    `docker attach /id/`

- Tổ hợp phím "Ctrl + P + Q": thoát terminal của container nhưng vẫn để container chạy
