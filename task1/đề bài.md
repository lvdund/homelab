### Write Dockerfile & Docker-compose file: Create 2 image

- Base from Golang image
- Base
- Golang app: 
    - use [Gin-gonic](https://github.com/gin-gonic/gin) to setup http server port 5000: method GET http://ip:5000/
    - each time call to this url: `countNumber += 1`
    - `countNumber` is save to mongoDB.
    - method GET http://ip:5000/count for checking `countNumber`.
- MongoDB: exposed at port 4000

** Note: push include your app!!!

==================================

### Viết Dockerfile và Docker-compose file: Tạo 2 image

- Base từ Golang image
- Base
- Golang app:
      - Dùng [Gin-gonic](https://github.com/gin-gonic/gin) để setup máy chú http cổng 5000: method GET http://ip:5000/
      - mỗi lần gọi đến url này: `countNumder += 1`
      - `countNumber` được lưu vào mongoDB
      - method Get http://ip:5000/count dể check `countNumber`
- MongoDB: ra ở cổng 4000

** Push cùng ứng dụng


====================================

# Cụ thể
## Tạo 2 images bằng cách sử dụng Docker, bao gồm một ứng dụng Golang và MongoDB để theo dõi số lần truy cập vào một endpoint HTTP. Cụ thể:

**Golang Image:**
- Ứng dụng này sử dụng framework Gin-Gonic để thiết lập một HTTP server hoạt động trên port 5000.
- Có 2 endpoint:
    - GET http://ip:5000/: Mỗi lần gọi vào đường dẫn này, biến countNumber sẽ tăng lên 1. Giá trị của biến countNumber này sẽ được lưu vào MongoDB.
    - GET http://ip:5000/count: Trả về giá trị hiện tại của countNumber, được lưu trong MongoDB.

**MongoDB Image:**
- Được sử dụng để lưu trữ giá trị của countNumber.
- MongoDB sẽ chạy và exposed (mở) port 4000 để ứng dụng Golang kết nối và lưu trữ dữ liệu.

**Yêu cầu khác:**
- Dockerfile cho ứng dụng Golang và MongoDB.
- Sử dụng docker-compose để khởi chạy cả 2 container này cùng lúc.
- Push mã nguồn bao gồm cả ứng dụng lên repository


23/10:
- Mongo: noSQL
- 