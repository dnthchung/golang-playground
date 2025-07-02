package main

import "fmt"

// TODO: Hoàn thành hàm này
func swap(a *int, b *int) {
    // Viết code của bạn ở đây để hoán đổi giá trị của a và b
	// temp := *a
	// *a = *b
	// *b = temp
	// Hoặc có thể viết ngắn gọn hơn:
	*a, *b = *b, *a
	// Cách này sử dụng tính năng hoán đổi của Go
}

func main() {
    x := 10
    y := 20

    fmt.Println("Trước khi hoán đổi:", x, y) // Kết quả mong đợi: 10 20

    swap(&x, &y)

    fmt.Println("Sau khi hoán đổi:", x, y) // Kết quả mong đợi: 20 10
}
