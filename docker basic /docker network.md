# Docker Network 
## Docker Network là một tính năng trong Docker cho phép các container giao tiếp với nhau và với các dịch vụ bên ngoài

- Hiểu đơn giản, mạng trong Docker giống như một hệ thống đường phố, nơi các container là những chiếc xe ô tô di chuyển, và mạng là các con đường mà chúng đi qua để giao tiếp với nhau.

**Lệnh cơ bản**
- Liệt kê các network đang có:
```
docker network ls
```
- Kiểm tra thông tin của network:
```
docker network inspect /name/
```
- Kiểm tra container đang kết nối với mạng nào:
```
docker inspect /name hoặc id/
```
- Tạo network:
```
docker network create /name tự đặt/
``` 

## Các loại mạng trong Docker

- Bridge Network: mặc định cho các container khi không chỉ định mạng cụ thể. Các container trong mạng cầu có thể giao tiếp với nhau thông qua địa chỉ IP hoặc tên container. 
 
- Host Network: Khi sử dụng mạng máy chủ, container sẽ chia sẻ mạng với máy chủ => container sẽ sử dụng địa chỉ IP của máy chủ và không có cách nào để cô lập mạng.  
 
- Null Network: Container sẽ không được kết nối đến bất kỳ mạng nào. Đây có thể là một lựa chọn khi muốn cô lập hoàn toàn một container. 

## Ví dụ cụ thể:
Giả sử có hai container: một cho ứng dụng web (nginx) và một cho cơ sở dữ liệu (MySQL). Muốn chúng có thể giao tiếp với nhau => tạo một mạng riêng cho chúng:

- Tạo mạng:
```
docker network create my_network
```
- Chạy container trên mạng đó:

```
docker run --name mydb --network my_network -e MYSQL_ROOT_PASSWORD=root -d mysql

docker run --name mynginx --network my_network -d nginx
```
