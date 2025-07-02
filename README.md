Dưới đây là **lịch học Go chi tiết theo giờ** trong 2 tuần (mỗi ngày 3 buổi), dựa trên roadmap bạn đã có. Bạn có thể điều chỉnh khung giờ cho phù hợp với lịch cá nhân.

---

## Tuần 1: Nền tảng và Kiến thức cơ bản

| Ngày         | 09:00 – 10:30                                                                         | 10:45 – 12:15                                                                         | 13:30 – 15:00                                                                  |
| ------------ | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| **Thứ 2**    | **Giới thiệu & Cài đặt**<br>– Cài Go, set GOPATH/GOROOT<br>– Cấu trúc thư mục project | **Basic Syntax**<br>– `package`, `import`, `func main()`<br>– Viết “Hello, Go!”       | **Thực hành**<br>– Biên dịch và chạy<br>– Chạy `go run`, `go build`            |
| **Thứ 3**    | **Biến & Kiểu dữ liệu**<br>– `var`, `:=`<br>– bool, string, int/float, rune, byte     | **Type Inference & Casting**<br>– Tự động suy luận<br>– Ép kiểu (`int()`, `string()`) | **Bài tập**<br>– Viết hàm tính diện tích hình chữ nhật với kiểu khác nhau      |
| **Thứ 4**    | **Điều kiện**<br>– `if`, `else`, `switch`                                             | **Vòng lặp**<br>– `for`, `for range`                                                  | **Bài tập**<br>– In số chẵn/lẻ 1–100, switch phân loại tuổi                    |
| **Thứ 5**    | **Hàm**<br>– Định nghĩa, tham số, trả về đa giá trị                                   | **Xử lý lỗi**<br>– `error` interface, `panic`/`recover`                               | **Thực hành**<br>– Viết hàm chia và xử lý chia cho 0                           |
| **Thứ 6**    | **Arrays & Slices**<br>– Khởi tạo, sử dụng `make()`                                   | **Maps**<br>– Tạo, truy xuất, xóa phần tử                                             | **Bài tập**<br>– Quản lý điểm số học sinh trong `map[string]float64`           |
| **Thứ 7**    | **Structs & Methods**<br>– Định nghĩa struct<br>– Receiver methods                    | **Packages**<br>– Tổ chức code, `import`/`export`                                     | **Mini-project**<br>– Xây module “user” với struct User và method `FullName()` |
| **Chủ nhật** | **Ôn tập & Quiz**<br>– Tóm tắt kiến thức Tuần 1                                       | **Code Challenge**<br>– 5 bài tập nhỏ                                                 | **Review**<br>– Chạy lại, debug, đọc tài liệu chính thức                       |

---

## Tuần 2: Đi sâu, Web & Microservices

| Ngày         | 09:00 – 10:30                                        | 10:45 – 12:15                                       | 13:30 – 15:00                                         |
| ------------ | ---------------------------------------------------- | --------------------------------------------------- | ----------------------------------------------------- |
| **Thứ 2**    | **Go Modules**<br>– `go mod init`<br>– `go mod tidy` | **JSON**<br>– `encoding/json`: Marshal/Unmarshal    | **Thực hành**<br>– Đọc ghi JSON từ file/HTTP          |
| **Thứ 3**    | **Goroutines**<br>– `go func()`, closure             | **Channels**<br>– buffered/unbuffered<br>– `select` | **Bài tập**<br>– Fan-in/fan-out pattern               |
| **Thứ 4**    | **Sync primitives**<br>– `sync.Mutex`, `WaitGroup`   | **Interface & Type Assertion**                      | **Mini-project**<br>– Hệ thống đếm truy cập đồng thời |
| **Thứ 5**    | **Web Framework**<br>Chọn Gin hoặc Echo              | **Routing & Middleware**                            | **Tạo REST API**<br>– CRUD đơn giản (to-do list)      |
| **Thứ 6**    | **ORM cơ bản (GORM)**<br>– Model, Migration          | **Kết nối DB**<br>– SQLite/Postgres                 | **Thực hành**<br>– Tạo, đọc, cập nhật, xóa (CRUD)     |
| **Thứ 7**    | **Logging**<br>– `log`/`slog`, Zerolog               | **HTTP Client**<br>– GRequests or `net/http`        | **Bài tập**<br>– Gửi request, log response            |
| **Chủ nhật** | **gRPC & Protobuf**<br>– Định nghĩa `.proto`         | **gRPC-Go**<br>– Server, Client                     | **Unit Testing & Benchmark**<br>– `testing` package   |

---

### Tips để tối ưu hiệu quả

1. **Giữ nhịp học**: Sau mỗi 45–50’ học lý thuyết, nghỉ 10’ rồi chuyển sang coding.
2. **Chơi thử code**: Luôn tự viết lại ví dụ, không chỉ copy–paste.
3. **Ghi chú**: Viết log trên GitHub Issues hoặc Notion để theo dõi tiến độ.
4. **Ôn lại**: Cuối tuần tổng kết, sửa bug, cải thiện code cũ.

Chúc bạn học hiệu quả và mau thành thạo Go! Nếu cần điều chỉnh lịch (thêm chủ đề, tăng thời gian…), cứ cho mình biết nhé.
