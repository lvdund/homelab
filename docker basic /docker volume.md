# Docker Volume
Là một cách để lưu trữ dữ liệu bên ngoài container 

## Đặc điểm Volume
- **Tách biệt dữ liệu và ứng dụng**: Khi chạy một ứng dụng trong Docker, dữ liệu của ứng dụng thường được lưu trữ trong container. Tuy nhiên, nếu xóa container, tất cả dữ liệu bên trong sẽ bị mất.  
=> Docker Volume cho phép lưu trữ dữ liệu ở một vị trí bên ngoài container, do đó, ngay cả khi container bị xóa, dữ liệu vẫn được giữ lại. 

- **Chia sẻ dữ liệu giữa các container**: Nếu có nhiều container cần truy cập cùng một dữ liệu, => sử dụng volume để chia sẻ dữ liệu giữa chúng. 

- **Quản lý đơn giản**: Docker cung cấp các lệnh đơn giản để tạo, liệt kê, xóa và quản lý các volume => dễ dàng quản lý dữ liệu mà không cần phải lo lắng về việc cấu hình hệ thống tệp hoặc đường dẫn. 

- **Dễ dàng sao lưu và phục hồi**: Vì volume lưu trữ dữ liệu ở vị trí bên ngoài container, việc sao lưu và phục hồi dữ liệu trở nên dễ dàng hơn.  

- **Hiệu suất tốt hơn**: Sử dụng volume thường cung cấp hiệu suất tốt hơn so với việc lưu trữ dữ liệu trong container, vì volume có thể được tối ưu hóa cho việc lưu trữ và truy cập dữ liệu. 


## Các lệnh cơ bản

- Tạo volume:
```
docker volume create <tên-volume>
```

- Liệt kê volume:
```
docker volume ps
```

- Gán volume:
```
docker run -v <tên-volume>:/đường/dẫn/trong/container <tên-image>
```
