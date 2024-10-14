# Docker build
- **docker build** là lệnh trong Docker được sử dụng để tạo ra một Docker image từ một tập hợp các chỉ dẫn được định nghĩa trong file gọi là Dockerfile.

## Các bước cơ bản của docker build:
- **Viết Dockerfile**: File này chứa các hướng dẫn, ví dụ như lấy hệ điều hành nền tảng (Ubuntu, Alpine,...), cài đặt phần mềm, thiết lập thư mục làm việc, v.v.
- **Chạy lệnh docker build**: Lệnh này sẽ đọc file Dockerfile và từng bước tạo ra image dựa trên các chỉ dẫn trong file.
- **Tạo image**: Sau khi hoàn tất, Docker sẽ tạo ra một Docker image mới => có thể lưu trữ hoặc sử dụng image này để tạo ra các container.
