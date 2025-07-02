package main

import "fmt"

func main() {
    // Dùng var
    var courseName string = "Go Fundamentals"
    fmt.Println("Khóa học:", courseName)

    // Dùng :=
    totalLessons := 14
    fmt.Println("Tổng số buổi học:", totalLessons)

    // Khai báo nhiều biến cùng lúc
    currentLesson, status := 2, "đang học"
    fmt.Println("Buổi học số", currentLesson, "có trạng thái:", status)

    // Kiểu bool
    isCompleted := false
    fmt.Println("Đã hoàn thành khóa học?", isCompleted)

    // Hằng số
    const welcomeMessage = "Chào mừng đến với Go!"
    fmt.Println(welcomeMessage)
}
