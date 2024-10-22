# Chọn image nền từ Golang
FROM golang:1.23 AS build

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép go.mod và go.sum
COPY go.mod go.sum ./
RUN go mod download

# Sao chép mã nguồn vào image
COPY . .

# Biên dịch ứng dụng
RUN go build -o main .

# Tạo image chính để chạy ứng dụng
FROM gcr.io/distroless/base

COPY --from=build /app/main /app/

# Chạy ứng dụng
CMD ["/app/main"]
