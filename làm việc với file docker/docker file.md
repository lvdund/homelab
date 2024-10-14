# Dockerfile 
- Dockerfile là một file dạng text chứa các lệnh viết ra để chỉ dẫn Docker xây dựng một image (hình ảnh) cho ứng dụng. Image này sau đó sẽ được dùng để tạo ra container, tức là môi trường độc lập nơi ứng dụng chạy.

## Trong Dockerfile có thể chỉ ra
- Hệ điều hành nào cần dùng (ví dụ: Ubuntu). 
- Những phần mềm nào cần cài đặt (ví dụ: Python, Node.js). 
- Các lệnh nào sẽ chạy khi container khởi động. 

## Cấu trúc Dockerfile cơ bản
```
FROM python:3.9-slim        # Mỗi Dockerfile phải bắt đầu với một lệnh FROM. Image nền tảng

WORKDIR /app                # Thiết lập thư mục làm việc bên trong container. Mọi lệnh tiếp theo trong Dockerfile sẽ thực thi trong thư mục này.

COPY requirements.txt .     # Sao chép các file cần thiết

RUN pip install -r requirements.txt  # Chạy các lệnh trong quá trình xây dựng image, ví dụ như cài đặt các gói phần mềm hoặc phụ thuộc cần thiết.

COPY . .                    # Sao chép toàn bộ mã nguồn

ENV FLASK_APP=app.py        # Thiết lập biến môi trường

EXPOSE 5000                 # Mở cổng 5000

CMD ["flask", "run", "--host=0.0.0.0"] # Chỉ định lệnh mặc định sẽ chạy khi container khởi động. Chỉ có một lệnh CMD trong mỗi Dockerfile (nếu có nhiều lệnh CMD, chỉ lệnh cuối cùng sẽ được sử dụng).
```

## Ví dụ Dockerfile cơ bản
- Sử dụng Ubuntu làm image nền tảng, và ứng dụng đơn giản: chạy một tệp script shell in ra dòng chữ "Hello, World!".
