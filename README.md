# GOLANG EXAMPLE

Đây là cấu trúc dự án tham khảo, không phải tiêu chuẩn, hãy áp dụng dựa theo tình huống của dự án.

- [GOLANG EXAMPLE](#golang-example)
  - [1. LUỒNG ĐI CỦA DỮ LIỆU](#1-luồng-đi-của-dữ-liệu)
  - [2. CẤU TRÚC DỰ ÁN](#2-cấu-trúc-dự-án)
  - [3. MỘT SỐ PACKAGE TIỆN ÍCH](#3-một-số-package-tiện-ích)
    - [3.1. Go validator](#31-go-validator)
    - [3.2. Automapper](#32-automapper)

## 1. LUỒNG ĐI CỦA DỮ LIỆU

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

### 3.2. [Automapper](https://pkg.go.dev/github.com/peteprogrammer/go-automapper)

**Sử dụng để tự động mapping giữa các model.**

Cài đặt:
```
go get github.com/peteprogrammer/go-automapper
```
