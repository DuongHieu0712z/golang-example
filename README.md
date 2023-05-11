# GOLANG EXAMPLE

Đây là code mẫu Web API Golang sử dụng Gin Framework và MongoDB.

*Lưu ý, cấu trúc dự án chỉ mang tính chất tham khảo.*

- [GOLANG EXAMPLE](#golang-example)
  - [1. KIẾN TRÚC LOGIC](#1-kiến-trúc-logic)
  - [2. CẤU TRÚC DỰ ÁN](#2-cấu-trúc-dự-án)
  - [3. MỘT SỐ PACKAGE TIỆN ÍCH](#3-một-số-package-tiện-ích)
    - [3.1. Go validator](#31-go-validator)
    - [3.2. devfeel/mapper](#32-devfeelmapper)

## 1. KIẾN TRÚC LOGIC

![business-logic](/document/image/business-logic.png)

## 2. CẤU TRÚC DỰ ÁN

- **`config`**: chứa các configuration của chương trình, lấy thông tin từ `.env`.
- **`db`**: chứa kết nối đến CSDL.
- **`model`**: chứa các model để lưu vào CSDL.
- **`repository`**: chứa các đối tượng thực hiện tương tác trực tiếp với CSDL như thêm, sửa, xóa,... Tham khảo thêm về *`repository pattern`*.
- **`uow` (`unit of work`)**: chứa đối tượng dùng để quản lý các repository và transaction. Tham khảo thêm về *`unit of work pattern`*.
- **`service`/`usecase`**: chứa các đối tượng thực hiện các business logic chính, kiểm tra nguồn dữ liệu trước khi đưa xuống `repository` và chuyển đổi dữ liệu trước khi trả về...
- **`controller`/`handler`**: chứa các đối tượng thực hiện xử lý request và trả về response. Không nên thực hiện xử lý business logic ở đây mà nên đưa xuống `service`/`usecase`.
- **`middleware`**: chứa các đối tượng `middleware`.
- **`routes`**: là nơi thực hiện mapping các `controller` với `route` tương ứng.
- **`form`**: chứa các đối tượng đầu vào, dùng để thực hiện từ binding request. Các đối tượng này thường được validate trước khi mapping sang `model`.
- **`dto` (`data to object`)**: chứa các đối tượng đầu ra, dùng để trả về response.
- **`common`**: chứa các đối tượng dùng chung.
- **`utils`**: chứa các hàm tiện ích.

## 3. MỘT SỐ PACKAGE TIỆN ÍCH

### 3.1. [Go validator](https://github.com/go-playground/validator)

**Sử dụng để tự động validate giá trị các trường trong struct.**

Cài đặt:

```
go get github.com/go-playground/validator/v10
```

### 3.2. [devfeel/mapper](https://pkg.go.dev/github.com/devfeel/mapper)

**Sử dụng để tự động mapping giữa các model.**

Cài đặt:

```
go get github.com/devfeel/mapper
```
